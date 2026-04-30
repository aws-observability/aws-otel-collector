// Package timberjack provides a rolling logger with size-based and time-based rotation.
//
// Timberjack is a simple, pluggable component for log rotation. It can rotate
// the active log file when any of the following occur:
//   - the file grows beyond MaxSize (size-based)
//   - the configured RotationInterval elapses (interval-based)
//   - a scheduled time is reached via RotateAt or RotateAtMinutes (clock-based)
//   - rotation is triggered explicitly via Rotate() (manual)
//
// Rotated files can optionally be compressed with gzip or zstd.
// Cleanup is handled automatically. Old log files are removed based on MaxBackups and MaxAge.
//
// Import:
//
//	import "github.com/DeRuina/timberjack"
//
// Timberjack works with any logger that writes to an io.Writer, including the
// standard library’s log package.
//
// Concurrency note: timberjack assumes a single process writes to the target
// files. Reusing the same Logger configuration across multiple processes on the
// same machine may lead to improper behavior.
//
// Source code: https://github.com/DeRuina/timberjack
package timberjack

import (
	"cmp"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"

	"github.com/klauspost/compress/zstd"
)

const (
	backupTimeFormat = "2006-01-02T15-04-05.000"
	compressSuffix   = ".gz"
	zstdSuffix       = ".zst"
	defaultMaxSize   = 100
	defaultFileMode  = 0o640
)

// ensure we always implement io.WriteCloser
var _ io.WriteCloser = (*Logger)(nil)

// safeClose is a generic function that safely closes a channel of any type.
// It prevents "panic: close of closed channel" and "panic: close of nil channel".
//
// The type parameter [T any] allows this function to work with channels of any element type,
// for example, chan int, chan string, chan struct{}, etc.
func safeClose[T any](ch chan T) {
	defer func() {
		recover()
	}()
	close(ch)
}

type rotateAt [2]int

// Logger is an io.WriteCloser that writes to the specified filename.
//
// Logger opens or creates the logfile on the first Write.
// If the file exists and is smaller than MaxSize megabytes, timberjack will open and append to that file.
// If the file's size exceeds MaxSize, or if the configured RotationInterval has elapsed since the last rotation,
// the file is closed, renamed with a timestamp, and a new logfile is created using the original filename.
//
// Thus, the filename you give Logger is always the "current" log file.
//
// Backups use the log file name given to Logger, in the form:
// `name-timestamp-<reason>.ext` where `name` is the filename without the extension,
// `timestamp` is the time of rotation formatted as `2006-01-02T15-04-05.000`,
// `reason` is "size" or "time" (Rotate/auto), or a custom tag (RotateWithReason), and `ext` is the original extension.
// For example, if your Logger.Filename is `/var/log/foo/server.log`, a backup created at 6:30pm on Nov 11 2016
// due to size would use the filename `/var/log/foo/server-2016-11-04T18-30-00.000-size.log`.
//
// # Cleaning Up Old Log Files
//
// Whenever a new logfile is created, old log files may be deleted based on MaxBackups and MaxAge.
// The most recent files (according to the timestamp) will be retained up to MaxBackups (or all files if MaxBackups is 0).
// Any files with a timestamp older than MaxAge days are deleted, regardless of MaxBackups.
// Note that the timestamp is the rotation time, not necessarily the last write time.
//
// If MaxBackups and MaxAge are both 0, no old log files will be deleted.
//
// timberjack assumes only a single process is writing to the log files at a time.
type Logger struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-timberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.) MaxBackups counts distinct rotation events (timestamps).
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Deprecated: use Compression instead ("none" | "gzip" | "zstd").
	Compress bool `json:"compress,omitempty" yaml:"compress,omitempty"`

	// Compression selects the algorithm. If empty, legacy Compress is used.
	// Allowed values: "none", "gzip", "zstd". Unknown => "none" (with a warning).
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// RotationInterval is the maximum duration between log rotations.
	// If the elapsed time since the last rotation exceeds this interval,
	// the log file is rotated, even if the file size has not reached MaxSize.
	// The minimum recommended value is 1 minute. If set to 0, time-based rotation is disabled.
	//
	// Example: RotationInterval = time.Hour * 24 will rotate logs daily.
	RotationInterval time.Duration `json:"rotationinterval" yaml:"rotationinterval"`

	// BackupTimeFormat defines the layout for the timestamp appended to rotated file names.
	// While other formats are allowed, it is recommended to follow the standard Go time layout
	// (https://pkg.go.dev/time#pkg-constants). Use the ValidateBackupTimeFormat() method to check
	// if the value is valid. It is recommended to call this method before using the Logger instance.
	//
	// WARNING: This field is assumed to be constant after initialization.
	// WARNING: If invalid value is supplied then default format `2006-01-02T15-04-05.000` will be used.
	//
	// Example:
	// BackupTimeFormat = `2006-01-02-15-04-05`
	// will generate rotated backup files in the format:
	// <logfilename>-2006-01-02-15-04-05-<rotationCriterion>-timberjack.log
	// where `rotationCriterion` could be `time` or `size`.
	BackupTimeFormat string `json:"backuptimeformat" yaml:"backuptimeformat"`

	// RotateAtMinutes defines specific minutes within an hour (0-59) to trigger a rotation.
	// For example, []int{0} for top of the hour, []int{0, 30} for top and half-past the hour.
	// Rotations are aligned to the clock minute (second 0).
	// This operates in addition to RotationInterval and MaxSize.
	// If multiple rotation conditions are met, the first one encountered typically triggers.
	RotateAtMinutes []int `json:"rotateAtMinutes" yaml:"rotateAtMinutes"`

	// RotateAt defines specific time within a day to trigger a rotation.
	// For example, []string{'00:00'} for midnight, []string{'00:00', '12:00'} for
	// midnight and midday.
	// Rotations are aligned to the clock minute (second 0).
	// This operates in addition to RotationInterval and MaxSize.
	// If multiple rotation conditions are met, the first one encountered typically triggers.
	RotateAt []string `json:"rotateAt" yaml:"rotateAt"`

	// AppendTimeAfterExt controls where the timestamp/reason go.
	// false (default):  <name>-<timestamp>-<reason>.log
	// true:             <name>.log-<timestamp>-<reason>
	AppendTimeAfterExt bool `json:"appendTimeAfterExt" yaml:"appendTimeAfterExt"`

	// FileMode sets the permissions to use when creating new log files.
	// It will be inherited by rotated files.
	// If zero, the default of 0o640 is used.
	//
	// Note that a local umask may alter the final permissions.
	// Also, on non-Linux systems this might not have the desired effect.
	FileMode os.FileMode `json:"filemode" yaml:"filemode"`

	// Internal fields
	size                   int64     // current size of the log file
	file                   *os.File  // current log file
	lastRotationTime       time.Time // records the last time a rotation happened (for interval/scheduled).
	logStartTime           time.Time // start time of the current logging period (used for backup filename timestamp).
	cfgOnce                sync.Once
	resolvedBackupLayout   string
	resolvedAppendAfterExt bool
	resolvedLocalTime      bool
	resolvedCompression    string

	mu sync.Mutex // ensures atomic writes and rotations

	// For mill goroutine (backups, compression cleanup)
	millCh        chan bool      // channel to signal the mill goroutine
	startMill     sync.Once      // ensures mill goroutine is started only once
	millWg        sync.WaitGroup // waits for the mill goroutine to finish
	millWGStarted bool

	// For scheduled rotation goroutine (RotateAt)
	startScheduledRotationOnce sync.Once      // ensures scheduled rotation goroutine is started only once
	scheduledRotationQuitCh    chan struct{}  // channel to signal the scheduled rotation goroutine to stop
	scheduledRotationWg        sync.WaitGroup // waits for the scheduled rotation goroutine to finish
	processedRotateAt          []rotateAt     // internal storage for sorted and validated RotateAt
	isClosed                   uint32

	// snapshots of globals to avoid races
	resolvedTimeNow func() time.Time
	resolvedStat    func(string) (os.FileInfo, error)
	resolvedRename  func(string, string) error
	resolvedRemove  func(string) error
}

var (
	// currentTime exists so it can be mocked out by tests.
	currentTime = time.Now

	// osStat exists so it can be mocked out by tests.
	osStat = os.Stat

	// megabyte is the conversion factor between MaxSize and bytes.  It is a
	// variable so tests can mock it out and not need to write megabytes of data
	// to disk.
	megabyte = 1024 * 1024

	osRename = os.Rename

	osRemove = os.Remove

	// empty BackupTimeFormatField
	ErrEmptyBackupTimeFormatField = errors.New("empty backupformat field")
)

func (l *Logger) resolveConfigLocked() {
	l.cfgOnce.Do(func() {
		// Resolve time format
		layout := l.BackupTimeFormat
		if layout == "" {
			layout = backupTimeFormat
		} else if err := l.ValidateBackupTimeFormat(); err != nil {
			fmt.Fprintf(os.Stderr,
				"timberjack: invalid BackupTimeFormat: %v — falling back to default format: %s\n",
				err, backupTimeFormat)
			layout = backupTimeFormat
		}
		l.resolvedBackupLayout = layout
		l.resolvedAppendAfterExt = l.AppendTimeAfterExt
		l.resolvedLocalTime = l.LocalTime

		// snapshot global funcs so tests can fiddle with the globals safely
		l.resolvedTimeNow = currentTime
		l.resolvedStat = osStat
		l.resolvedRename = osRename
		l.resolvedRemove = osRemove

		// Freeze compression (prevents races if toggled later)
		l.resolvedCompression = l.effectiveCompression()
	})
}

// Write implements io.Writer.
// It writes the provided bytes to the current log file.
// If the log file exceeds MaxSize after writing, or if the configured RotationInterval has elapsed
// since the last rotation, or if a scheduled rotation time (RotateAtMinutes) has been reached,
// the file is closed, renamed to include a timestamp, and a new log file is created
// using the original filename.
// If the size of a single write exceeds MaxSize, the write is rejected and an error is returned.
func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Ensure configuration is resolved once
	l.resolveConfigLocked()

	// Handle writes to a closed logger.
	if atomic.LoadUint32(&l.isClosed) == 1 {
		// The logger is closed. To ensure the write succeeds, we perform a
		// single open-write-close cycle. This does not perform rotation
		// and does not restart the background goroutines. l.file remains nil.
		file, openErr := os.OpenFile(l.filename(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if openErr != nil {
			return 0, fmt.Errorf("timberjack: write on closed logger failed to open file: %w", openErr)
		}

		n, writeErr := file.Write(p)

		closeErr := file.Close()

		if writeErr != nil {
			return n, writeErr
		}
		return n, closeErr
	}

	// Ensure the scheduled-rotation goroutine is running (if you've still got one).
	l.ensureScheduledRotationLoopRunning()

	// Snapshot
	now := l.resolvedTimeNow().In(l.location())

	writeLen := int64(len(p))
	if writeLen > l.max() {
		return 0, fmt.Errorf("write length %d exceeds maximum file size %d", writeLen, l.max())
	}

	// Open (or create) the file on first write.
	if l.file == nil {
		if err = l.openExistingOrNew(len(p)); err != nil {
			return 0, err
		}
		if l.lastRotationTime.IsZero() {
			// Initialize to 'now' so interval/minute checks start from here.
			l.lastRotationTime = now
		}
	}

	// 1) Interval-based rotation
	if l.RotationInterval > 0 && now.Sub(l.lastRotationTime) >= l.RotationInterval {
		if err := l.rotate("time"); err != nil {
			return 0, fmt.Errorf("interval rotation failed: %w", err)
		}
		l.lastRotationTime = now
	}

	// 2) Scheduled time based rotation (RotateAt)
	if len(l.processedRotateAt) > 0 {
		for _, m := range l.processedRotateAt {
			mark := time.Date(now.Year(), now.Month(), now.Day(),
				m[0], m[1], 0, 0, l.location())
			// If we've crossed that mark since the last rotation, fire one rotation.
			if l.lastRotationTime.Before(mark) && (mark.Before(now) || mark.Equal(now)) {
				if err := l.rotate("time"); err != nil {
					return 0, fmt.Errorf("scheduled-minute rotation failed: %w", err)
				}
				// Record the logical mark—so we don’t rerun until next slot.
				l.lastRotationTime = mark
				break
			}
		}
	}

	// 3) Size-based rotation
	if l.size+writeLen > l.max() {
		if err := l.rotate("size"); err != nil {
			return 0, fmt.Errorf("size rotation failed: %w", err)
		}
		// Note: we leave lastRotationTime untouched for size rotations.
	}

	// Finally, write the bytes and update size.
	n, err = l.file.Write(p)
	l.size += int64(n)
	return n, err
}

// ValidateBackupTimeFormat checks if the configured BackupTimeFormat is a valid time layout.
// While other formats are allowed, it is recommended to follow the standard time layout
// rules as defined here: https://pkg.go.dev/time#pkg-constants
//
// WARNING: Assumes that BackupTimeFormat value remains constant after initialization.
func (l *Logger) ValidateBackupTimeFormat() error {
	if len(l.BackupTimeFormat) == 0 {
		return ErrEmptyBackupTimeFormatField
	}
	// 2025-05-22 23:41:59.987654321 +0000 UTC
	now := time.Date(2025, 5, 22, 23, 41, 59, 987_654_321, time.UTC)

	layoutPrecision := countDigitsAfterDot(l.BackupTimeFormat)

	now, err := truncateFractional(now, layoutPrecision)
	if err != nil {
		return err
	}
	formatted := now.Format(l.BackupTimeFormat)
	parsedT, err := time.Parse(l.BackupTimeFormat, formatted)
	if err != nil {
		return fmt.Errorf("invalid BackupTimeFormat: %w", err)
	}
	if !parsedT.Equal(now) {
		return errors.New("invalid BackupTimeFormat: time.Time parsed from the format does not match the time.Time supplied")
	}

	return nil
}

// location returns the time.Location (UTC or Local) to use for timestamps in backup filenames.
func (l *Logger) location() *time.Location {
	l.resolveConfigLocked()
	if l.resolvedLocalTime {
		return time.Local
	}
	return time.UTC
}

func parseTime(s string) (*rotateAt, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, errors.New("invalid time")
	}
	hs, ms := parts[0], parts[1]
	h, err := strconv.Atoi(hs)
	if err != nil {
		return nil, err
	}
	m, err := strconv.Atoi(ms)
	if err != nil {
		return nil, err
	}

	if m < 0 || m > 59 || h < 0 || h > 23 {
		return nil, errors.New("invalid time")
	}

	return &rotateAt{h, m}, nil
}

func compareTime(a, b rotateAt) int {
	h1, m1 := a[0], a[1]
	h2, m2 := b[0], b[1]
	if h1 == h2 {
		return cmp.Compare(m1, m2)
	}
	return cmp.Compare(h1, h2)
}

// ensureScheduledRotationLoopRunning starts the scheduled rotation goroutine if RotateAtMinutes is configured
// and the goroutine is not already running.
func (l *Logger) ensureScheduledRotationLoopRunning() {
	if len(l.RotateAtMinutes)+len(l.RotateAt) == 0 {
		return // No scheduled rotations configured
	}

	l.startScheduledRotationOnce.Do(func() {
		var processedRotateAt []rotateAt
		for _, m := range l.RotateAtMinutes {
			if m < 0 || m > 59 {
				fmt.Fprintf(os.Stderr, "timberjack: [%d] No valid minute specified for RotateAtMinutes.\n", m)
				continue
			}
			for h := 0; h < 24; h++ {
				processedRotateAt = append(processedRotateAt, rotateAt{h, m})
			}
		}

		for _, t := range l.RotateAt {
			r, err := parseTime(t)
			if err != nil {
				fmt.Fprintf(os.Stderr, "timberjack: [%s] No valid time specified for RotateAt.\n", t)
				continue
			}
			processedRotateAt = append(processedRotateAt, *r)
		}

		if len(processedRotateAt) == 0 {
			// Optionally log that no valid minutes were found, preventing goroutine start
			// fmt.Fprintf(os.Stderr, "timberjack: [%s] No valid minutes specified for RotateAtMinutes.\n", l.Filename)
			return
		}

		// Sort for predictable order in calculating next rotation
		slices.SortFunc(processedRotateAt, compareTime)
		processedRotateAt = slices.CompactFunc(processedRotateAt, func(a, b rotateAt) bool {
			return compareTime(a, b) == 0
		})

		l.processedRotateAt = processedRotateAt
		quitCh := make(chan struct{})
		l.scheduledRotationQuitCh = quitCh

		// Snapshot immutable inputs for the goroutine to avoid reading fields later.
		slots := l.processedRotateAt
		loc := l.location()
		nowFn := l.resolvedTimeNow

		l.scheduledRotationWg.Add(1)
		go l.runScheduledRotations(quitCh, slots, loc, nowFn)
	})
}

// runScheduledRotations is the main loop for handling rotations at specific minute marks
// as defined in RotateAtMinutes. It runs in a separate goroutine.
func (l *Logger) runScheduledRotations(quit <-chan struct{}, slots []rotateAt, loc *time.Location, nowFn func() time.Time) {
	defer l.scheduledRotationWg.Done()

	// No slots to process, exit immediately.
	if len(slots) == 0 {
		return
	}

	timer := time.NewTimer(0) // Timer will be reset with the correct duration in the loop
	if !timer.Stop() {
		// Drain the channel if the timer fired prematurely (e.g., duration was 0 on first NewTimer)
		select {
		case <-timer.C:
		default:
		}
	}

	for {
		now := nowFn()
		nowInLocation := now.In(loc)
		nextRotationAbsoluteTime := time.Time{}
		foundNextSlot := false

	determineNextSlot:
		// Calculate the next rotation time based on the current time and processedRotateAt.
		// Iterate through the current hour, then subsequent hours (up to 24h ahead for robustness
		// against system sleep or large clock jumps).
		for hourOffset := 0; hourOffset <= 24; hourOffset++ {
			// Base time for the hour we are checking (e.g., if now is 10:35, current hour base is 10:00)
			hourToCheck := time.Date(nowInLocation.Year(), nowInLocation.Month(), nowInLocation.Day(), nowInLocation.Hour(), 0, 0, 0, loc).Add(time.Duration(hourOffset) * time.Hour)

			for _, mark := range slots {
				candidateTime := time.Date(hourToCheck.Year(), hourToCheck.Month(), hourToCheck.Day(), mark[0], mark[1], 0, 0, loc)

				if candidateTime.After(now) { // Found the earliest future slot
					nextRotationAbsoluteTime = candidateTime
					foundNextSlot = true
					break determineNextSlot // Exit both loops
				}
			}
		}

		if !foundNextSlot {
			// This should ideally not happen if processedRotateAt is valid and non-empty.
			// Could occur if currentTime() is unreliable or jumps massively backward.
			// Log an error and retry calculation after a fallback delay.
			fmt.Fprintf(os.Stderr, "timberjack: [%s] Could not determine next scheduled rotation time for %v with marks %v. Retrying calculation in 1 minute.\n", l.Filename, nowInLocation, slots)
			select {
			case <-time.After(time.Minute): // Wait a bit before retrying calculation
				continue // Restart the outer loop to recalculate
			case <-quit:
				return
			}
		}

		sleepDuration := nextRotationAbsoluteTime.Sub(now)
		timer.Reset(sleepDuration)

		select {
		case <-timer.C: // Timer fired, it's time for a scheduled rotation
			l.mu.Lock()
			// Only rotate if the last rotation time was before this specific scheduled mark.
			// This prevents redundant rotations if another rotation (e.g., size/interval) happened
			// very close to, but just before or at, this scheduled time for the same mark.
			if l.lastRotationTime.Before(nextRotationAbsoluteTime) {
				if err := l.rotate("time"); err != nil { // Scheduled rotations are "time" based for filename
					fmt.Fprintf(os.Stderr, "timberjack: [%s] scheduled rotation failed: %v\n", l.Filename, err)
				} else {
					l.lastRotationTime = nowFn()
				}
			}
			l.mu.Unlock()
		// Loop will continue and recalculate the next slot from the new "now"

		case <-quit: // Signal to quit from Close()
			if !timer.Stop() {
				// If Stop() returns false, the timer has already fired or been stopped.
				// If it fired, its channel might have a value, so drain it.
				select {
				case <-timer.C:
				default:
				}
			}
			return // Exit goroutine
		}
	}
}

// Close implements io.Closer, and closes the current logfile.
// It also signals any running goroutines (like scheduled rotation or mill) to stop.
func (l *Logger) Close() error {
	l.mu.Lock()

	if atomic.LoadUint32(&l.isClosed) == 1 {
		l.mu.Unlock()
		return nil
	}
	atomic.StoreUint32(&l.isClosed, 1)

	// Stop the scheduled rotation goroutine
	var quitCh chan struct{}
	if l.scheduledRotationQuitCh != nil {
		quitCh = l.scheduledRotationQuitCh
		l.scheduledRotationQuitCh = nil // clear under lock
	}

	// Stop the mill goroutine (doesn't use l.mu, no deadlock risk)
	var millCh chan bool
	if l.millCh != nil {
		millCh = l.millCh
		// don't nil it; startMill.Do prevents restarts
	}

	// Close file under lock (safe and quick)
	err := l.closeFile()

	l.mu.Unlock()

	// Now signal and wait outside the lock
	if quitCh != nil {
		safeClose(quitCh)
		l.scheduledRotationWg.Wait()
	}
	if millCh != nil {
		safeClose(millCh)
		l.millWg.Wait()
	}

	return err
}

// Sync flushes buffered data to the underlying file, satisfying the
// zapcore.WriteSyncer interface. It is a no-op when no file is currently open
// (e.g. before the first write or after Close).
func (l *Logger) Sync() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if atomic.LoadUint32(&l.isClosed) == 1 {
		return nil
	}
	if l.file == nil {
		return nil
	}
	return l.file.Sync()
}

// closeFile closes the file if it is open. This is an internal method.
// It expects l.mu to be held.
func (l *Logger) closeFile() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil // Set to nil to indicate it's closed.
	return err
}

// Rotate forces an immediate rotation using the legacy auto-reason logic.
// (empty reason => "time" if an interval rotation is due, otherwise "size")
func (l *Logger) Rotate() error {
	return l.RotateWithReason("")
}

// rotate closes the current file, moves it aside with a timestamp in the name,
// (if it exists), opens a new file with the original filename, and then runs
// post-rotation processing and removal (mill).
// It expects l.mu to be held by the caller.
// Takes an explicit reason for the rotation which is used in the backup filename.
func (l *Logger) rotate(reason string) error {
	if err := l.closeFile(); err != nil {
		return err
	}
	// Pass the determined reason to openNew so it's used in the backup filename
	if err := l.openNew(reason); err != nil {
		return err
	}
	l.mill()
	return nil
}

// RotateWithReason forces a rotation immediately and tags the backup filename
// with the provided reason (after sanitization). If the sanitized reason is
// empty, it falls back to the default behavior used by Rotate(): "time" if an
// interval rotation is due, otherwise "size".
//
// After a successful rotation, lastRotationTime is updated to the current time
// so that a subsequent Write() does not trigger a duplicate interval-based or
// scheduled rotation.
func (l *Logger) RotateWithReason(reason string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if atomic.LoadUint32(&l.isClosed) == 1 {
		return errors.New("logger closed")
	}

	r := sanitizeReason(reason)
	if r == "" {
		// keep legacy Rotate() semantics
		r = "size"
		if l.shouldTimeRotate() {
			r = "time"
		}
	}

	if err := l.rotate(r); err != nil {
		return err
	}
	l.lastRotationTime = l.resolvedTimeNow()
	return nil
}

func backupNameWithResolved(name string, local bool, reason string, t time.Time, layout string, afterExt bool) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]

	loc := time.UTC
	if local {
		loc = time.Local
	}
	ts := t.In(loc).Format(layout)

	if afterExt {
		// <name><ext>-<ts>-<reason>
		return filepath.Join(dir, fmt.Sprintf("%s%s-%s-%s", prefix, ext, ts, reason))
	}
	// <name>-<ts>-<reason><ext>
	return filepath.Join(dir, fmt.Sprintf("%s-%s-%s%s", prefix, ts, reason, ext))
}

// openNew creates a new log file for writing.
// If an old log file already exists, it is moved aside by renaming it with a timestamp.
// This method assumes that l.mu is held and the old file (if any) has already been closed.
// The reasonForBackup parameter is used in the backup filename.
func (l *Logger) openNew(reasonForBackup string) error {
	l.resolveConfigLocked() // no-op after first time

	if err := os.MkdirAll(l.dir(), 0o755); err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)
	}

	name := l.filename()

	finalMode := l.FileMode
	if finalMode == 0 {
		finalMode = os.FileMode(defaultFileMode)
	}

	var oldInfo os.FileInfo
	info, err := l.resolvedStat(name)
	if err == nil {
		oldInfo = info
		// Only use the existing file's mode when no explicit FileMode is configured.
		if l.FileMode == 0 {
			finalMode = oldInfo.Mode()
		}

		rotationTimeForBackup := l.resolvedTimeNow()

		// Build the rotated name from the immutable snapshot (no public field writes).
		newname := backupNameWithResolved(
			name,
			l.resolvedLocalTime,
			reasonForBackup,
			rotationTimeForBackup,
			l.resolvedBackupLayout,
			l.resolvedAppendAfterExt,
		)

		if errRename := l.resolvedRename(name, newname); errRename != nil {
			return fmt.Errorf("can't rename log file: %s", errRename)
		}
		l.logStartTime = rotationTimeForBackup
	} else if os.IsNotExist(err) {
		l.logStartTime = l.resolvedTimeNow()
		oldInfo = nil
	} else {
		return fmt.Errorf("failed to stat log file %s: %w", name, err)
	}

	// Create and open the new log file at path `name`.
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, finalMode)
	if err != nil {
		return fmt.Errorf("can't open new logfile %s: %s", name, err)
	}
	// Apply the exact mode via chmod so the process umask cannot mask bits.
	if err := os.Chmod(name, finalMode); err != nil {
		closeErr := f.Close()
		if closeErr != nil {
			return fmt.Errorf("can't set mode on new logfile %s: %s (also failed to close: %v)", name, err, closeErr)
		}
		return fmt.Errorf("can't set mode on new logfile %s: %s", name, err)
	}
	l.file = f
	l.size = 0

	// Try to chown the new file to match the old file's owner/group (if there was an old file).
	if oldInfo != nil {
		if errChown := chown(name, oldInfo); errChown != nil {
			fmt.Fprintf(os.Stderr, "timberjack: [%s] failed to chown new log file %s: %v\n", l.Filename, name, errChown)
		}
	}

	return nil
}

// shouldTimeRotate checks if the time-based rotation interval has elapsed
// since the last rotation. This is used for RotationInterval logic.
func (l *Logger) shouldTimeRotate() bool {
	l.resolveConfigLocked()
	if l.RotationInterval == 0 { // Time-based rotation (interval) is disabled
		return false
	}
	// If lastRotationTime is zero (e.g., logger just started, no writes/rotations yet),
	// then it's not yet time for an interval-based rotation.
	if l.lastRotationTime.IsZero() {
		return false
	}
	return l.resolvedTimeNow().Sub(l.lastRotationTime) >= l.RotationInterval
}

// backupName creates a new backup filename by inserting a timestamp and a rotation reason
// ("time" or "size") between the filename prefix and the extension.
// It uses the local time if requested (otherwise UTC).
func backupName(name string, local bool, reason string, t time.Time, fileTimeFormat string, appendTimeAfterExt bool) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]

	currentLoc := time.UTC
	if local {
		currentLoc = time.Local
	}
	// Format the timestamp for the backup file.
	timestamp := t.In(currentLoc).Format(fileTimeFormat)

	if appendTimeAfterExt {
		// <name><ext>-<ts>-<reason>
		// e.g. httpd.log-2025-01-01T00-00-00.000-size
		return filepath.Join(dir, fmt.Sprintf("%s%s-%s-%s", prefix, ext, timestamp, reason))
	}

	// default: <name>-<ts>-<reason><ext>
	return filepath.Join(dir, fmt.Sprintf("%s-%s-%s%s", prefix, timestamp, reason, ext))
}

// openExistingOrNew opens the existing logfile if it exists and the current write
// would not cause it to exceed MaxSize. If the file does not exist, or if writing
// would exceed MaxSize, the current file is rotated (if it exists) and a new logfile is created.
// It expects l.mu to be held by the caller.
func (l *Logger) openExistingOrNew(writeLen int) error {
	l.resolveConfigLocked()
	l.mill() // Perform house-keeping for old logs (compression, deletion) first.

	filename := l.filename()
	info, err := l.resolvedStat(filename)
	if os.IsNotExist(err) {
		// File doesn't exist, so openNew is creating a new file.
		// The 'reason' passed to openNew here ("initial") won't affect a backup filename
		// as no backup (renaming) is happening.
		return l.openNew("initial")
	}
	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	// Check if rotation is needed due to size before opening/appending.
	if info.Size()+int64(writeLen) >= l.max() {
		return l.rotate("size") // This rotation is explicitly due to "size"
	}

	// Open existing file for appending.
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		// If opening existing fails (e.g., permissions, corruption), try to create a new one.
		return l.openNew("initial") // Fallback if append fails
	}
	l.file = file
	l.size = info.Size()
	// Note: l.logStartTime is NOT updated here if we successfully open an existing file without rotating.
	// It retains its value from when this current log segment was created (by a previous openNew).
	// l.lastRotationTime is also NOT updated here; it's handled by rotation trigger logic.
	return nil
}

// filename returns the current log filename, using the configured Filename,
// or a default based on the process name if Filename is empty.
func (l *Logger) filename() string {
	if l.Filename != "" {
		return l.Filename
	}
	name := filepath.Base(os.Args[0]) + "-timberjack.log"
	return filepath.Join(os.TempDir(), name)
}

// millRunOnce performs one cycle of compression and removal of old log files.
// If compression is enabled, uncompressed backups are compressed using gzip.
// Old backup files are deleted to enforce MaxBackups and MaxAge limits.
func (l *Logger) millRunOnce() error {
	l.resolveConfigLocked()
	comp := l.resolvedCompression
	if l.MaxBackups == 0 && l.MaxAge == 0 && comp == "none" {
		return nil // Nothing to do if all cleanup options are disabled.
	}

	now := l.resolvedTimeNow()

	files, err := l.oldLogFiles() // Gets LogInfo structs, sorted newest first by timestamp
	if err != nil {
		return err
	}

	filesToProcess := files     // Start with all found old log files
	var filesToRemove []logInfo // Accumulates files to be deleted

	// MaxBackups filtering: Keep files belonging to the MaxBackups newest distinct timestamps
	if l.MaxBackups > 0 {
		uniqueTimestamps := make([]time.Time, 0)
		timestampMap := make(map[time.Time]bool)
		for _, f := range filesToProcess { // filesToProcess is sorted newest first
			if !timestampMap[f.timestamp] {
				timestampMap[f.timestamp] = true
				uniqueTimestamps = append(uniqueTimestamps, f.timestamp)
			}
		}

		if len(uniqueTimestamps) > l.MaxBackups {
			// Determine the set of timestamps to keep (the MaxBackups newest ones)
			keptTimestampsSet := make(map[time.Time]bool)
			for i := 0; i < l.MaxBackups; i++ {
				keptTimestampsSet[uniqueTimestamps[i]] = true
			}

			var filteredFiles []logInfo // Files that pass this MaxBackups filter
			for _, f := range filesToProcess {
				if keptTimestampsSet[f.timestamp] {
					filteredFiles = append(filteredFiles, f)
				} else {
					filesToRemove = append(filesToRemove, f) // Mark for removal
				}
			}
			filesToProcess = filteredFiles // Update filesToProcess for subsequent filters
		}
		// If len(uniqueTimestamps) <= l.MaxBackups, all files pass this MaxBackups filter.
	}

	// MaxAge filtering (operates on files that passed MaxBackups filter)
	if l.MaxAge > 0 {
		diff := time.Duration(int64(24*time.Hour) * int64(l.MaxAge))
		cutoff := now.Add(-diff)
		var filteredFiles []logInfo // Files that pass this MaxAge filter
		for _, f := range filesToProcess {
			if f.timestamp.Before(cutoff) {
				// avoid duplicates
				isAlreadyMarked := false
				for _, rmf := range filesToRemove {
					if rmf.Name() == f.Name() {
						isAlreadyMarked = true
						break
					}
				}
				if !isAlreadyMarked {
					filesToRemove = append(filesToRemove, f)
				}
			} else {
				filteredFiles = append(filteredFiles, f)
			}
		}
		filesToProcess = filteredFiles
	}

	// Compression task identification (operates on files that passed MaxBackups and MaxAge)
	var filesToCompress []logInfo
	if comp != "none" {
		for _, f := range filesToProcess {
			name := f.Name()
			if strings.HasSuffix(name, compressSuffix) || strings.HasSuffix(name, zstdSuffix) {
				continue // already compressed
			}
			isMarked := false
			for _, rmf := range filesToRemove {
				if rmf.Name() == name {
					isMarked = true
					break
				}
			}
			if !isMarked {
				filesToCompress = append(filesToCompress, f)
			}
		}
	}

	// Execute removals (ensure unique removals)
	finalUniqueRemovals := make(map[string]logInfo)
	for _, f := range filesToRemove {
		finalUniqueRemovals[f.Name()] = f
	}
	for _, f := range finalUniqueRemovals {
		errRemove := l.resolvedRemove(filepath.Join(l.dir(), f.Name()))
		if errRemove != nil && !os.IsNotExist(errRemove) {
			fmt.Fprintf(os.Stderr, "timberjack: [%s] failed to remove old log file %s: %v\n", l.Filename, f.Name(), errRemove)
		}
	}

	// Execute compressions (suffix also comes from snapshot via compressedSuffix)
	suffix := l.compressedSuffix()
	for _, f := range filesToCompress {
		fn := filepath.Join(l.dir(), f.Name())
		if errCompress := l.compressLogFile(fn, fn+suffix); errCompress != nil {
			fmt.Fprintf(os.Stderr, "timberjack: [%s] failed to compress log file %s: %v\n", l.Filename, f.Name(), errCompress)
		}
	}
	return nil
}

// millRun runs in a goroutine to manage post-rotation compression and removal
// of old log files. It listens on millCh for signals to run millRunOnce.
func (l *Logger) millRun() {
	if l.millWGStarted {
		defer l.millWg.Done()
	}
	ch := l.millCh
	for range ch {
		_ = l.millRunOnce()
	}
}

// mill performs post-rotation compression and removal of stale log files,
// starting the mill goroutine if necessary and sending a signal to it.
func (l *Logger) mill() {
	if atomic.LoadUint32(&l.isClosed) == 1 {
		return
	}
	l.startMill.Do(func() {
		l.millWGStarted = true
		l.millCh = make(chan bool, 1)
		l.millWg.Add(1)
		go l.millRun()
	})
	select {
	case l.millCh <- true:
	default:
	}
}

// oldLogFiles returns the list of backup log files stored in the same
// directory as the current log file, sorted by their embedded timestamp (newest first).
func (l *Logger) oldLogFiles() ([]logInfo, error) {
	entries, err := os.ReadDir(l.dir()) // ReadDir is generally preferred over ReadFile for directory listings
	if err != nil {
		return nil, fmt.Errorf("can't read log file directory: %s", err)
	}
	var logFiles []logInfo

	prefix, ext := l.prefixAndExt() // Get prefix like "filename-" and original extension like ".log"

	for _, e := range entries {
		if e.IsDir() { // Skip directories
			continue
		}
		name := e.Name()
		info, errInfo := e.Info() // Get FileInfo for modification time and other details
		if errInfo != nil {
			// fmt.Fprintf(os.Stderr, "timberjack: failed to get FileInfo for %s: %v\n", name, errInfo)
			continue // Skip files we can't stat
		}

		// Attempt to parse timestamp from filename (e.g., from "filename-timestamp-reason.log")
		if t, errTime := l.timeFromName(name, prefix, ext); errTime == nil {
			logFiles = append(logFiles, logInfo{t, info})
			continue
		}
		// Attempt to parse timestamp from compressed gzip filename (e.g., from "filename-timestamp-reason.log.gz")
		if t, errTime := l.timeFromName(name, prefix, ext+compressSuffix); errTime == nil {
			logFiles = append(logFiles, logInfo{t, info})
			continue
		}
		// Attempt to parse timestamp from compressed zstd filename (e.g., from "filename-timestamp-reason.log.zst")
		if t, errTime := l.timeFromName(name, prefix, ext+zstdSuffix); errTime == nil {
			logFiles = append(logFiles, logInfo{t, info})
			continue
		}
		// Files that don't match the expected backup pattern are ignored.
	}

	sort.Sort(byFormatTime(logFiles)) // Sorts newest first based on parsed timestamp
	return logFiles, nil
}

// timeFromName extracts the formatted timestamp from the backup filename.
// It expects filenames like "prefix-YYYY-MM-DDTHH-MM-SS.mmm-reason.ext" or "prefix.ext-YYYY-MM-DDTHH-MM-SS.mmm-reason[.gz]"
func (l *Logger) timeFromName(filename, prefix, ext string) (time.Time, error) {
	layout := l.resolvedBackupLayout
	if layout == "" {
		// defensive default if called very early
		layout = backupTimeFormat
	}
	loc := time.UTC
	if l.resolvedLocalTime {
		loc = time.Local
	}

	if !l.resolvedAppendAfterExt {

		// Keep legacy behavior for error messages to satisfy existing tests
		if !strings.HasPrefix(filename, prefix) {
			return time.Time{}, errors.New("mismatched prefix")
		}
		if !strings.HasSuffix(filename, ext) {
			return time.Time{}, errors.New("mismatched extension")
		}
		// "<prefix><timestamp>-<reason><ext>"
		trimmed := filename[len(prefix) : len(filename)-len(ext)]
		lastHyphenIdx := strings.LastIndex(trimmed, "-")
		if lastHyphenIdx == -1 {
			return time.Time{}, fmt.Errorf("malformed backup filename: missing reason separator in %q", trimmed)
		}
		ts := trimmed[:lastHyphenIdx]
		return time.ParseInLocation(layout, ts, loc)
	}

	// After-ext parsing:
	// base is "<name><ext>" (e.g., "foo.log")
	base := prefix[:len(prefix)-1] + ext

	// Allow optional trailing compression suffix (".gz" or ".zst")
	nameNoComp := trimCompressionSuffix(filename)

	// nameNoComp must start with "<base>-"
	if !strings.HasPrefix(nameNoComp, base+"-") {
		return time.Time{}, fmt.Errorf("malformed backup filename: %q", filename)
	}

	// nameNoComp = "<base>-<timestamp>-<reason>"
	trimmed := nameNoComp[len(base)+1:]

	lastHyphenIdx := strings.LastIndex(trimmed, "-")
	if lastHyphenIdx == -1 {
		return time.Time{}, fmt.Errorf("malformed backup filename: %q", filename)
	}
	ts := trimmed[:lastHyphenIdx]
	return time.ParseInLocation(layout, ts, loc)
}

// max returns the maximum size in bytes of log files before rolling.
func (l *Logger) max() int64 {
	if l.MaxSize == 0 { // If MaxSize is 0, use default.
		return int64(defaultMaxSize * megabyte)
	}
	return int64(l.MaxSize) * int64(megabyte)
}

// dir returns the directory for the current filename.
func (l *Logger) dir() string {
	return filepath.Dir(l.filename())
}

// prefixAndExt returns the filename part (up to the extension, with a trailing dash for backups)
// and extension part from the Logger's filename.
// e.g., for "foo.log", returns "foo-", ".log"
func (l *Logger) prefixAndExt() (prefix, ext string) {
	filename := filepath.Base(l.filename())
	ext = filepath.Ext(filename)
	prefix = filename[:len(filename)-len(ext)] + "-" // Add dash as backup filenames include it after original prefix
	return prefix, ext
}

// countDigitsAfterDot returns the number of consecutive digit characters
// immediately following the first '.' in the input.
// It skips all characters before the '.' and stops counting at the first non-digit
// character after the '.'.

// Example: `prefix.0012304123suffix` would return 10
// Example: `prefix.0012304_middle_123_suffix` would return 7
func countDigitsAfterDot(layout string) int {
	for i, ch := range layout {
		if ch == '.' {
			count := 0
			for _, c := range layout[i+1:] {
				if unicode.IsDigit(c) {
					count++
				} else {
					break
				}
			}
			return count
		}
	}
	return 0 // no '.' found or no digits after dot
}

// truncateFractional truncates time t to n fractional digits of seconds.
// n=0 → truncate to seconds, n=3 → milliseconds, n=6 → microseconds, etc.
func truncateFractional(t time.Time, n int) (time.Time, error) {
	if n < 0 || n > 9 {
		return time.Time{}, fmt.Errorf("unsupported fractional precision: %d", n)
	}

	// number of nanoseconds to keep
	factor := math.Pow10(9 - n) // e.g. for n=3, factor=10^(9-3)=1,000,000

	nanos := t.Nanosecond()
	truncatedNanos := int((int64(nanos) / int64(factor)) * int64(factor))

	return time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
		truncatedNanos,
		t.Location(),
	), nil
}

// compressLogFile compresses the given source log file (src) to a destination file (dst),
// removing the source file if compression is successful.
//
// Compression is written to a temporary file (dst+".tmp") first, then atomically
// renamed to dst. This ensures dst never appears partially-written on disk.
func (l *Logger) compressLogFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source log file %s for compression: %v", src, err)
	}
	defer srcFile.Close()

	srcInfo, err := l.resolvedStat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source log file %s: %v", src, err)
	}

	// Write to a temp file so that dst only appears once fully written (atomic publish).
	tmpDst := dst + ".tmp"
	dstFile, err := os.OpenFile(tmpDst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to open destination compressed log file %s: %v", tmpDst, err)
	}
	// No `defer dstFile.Close()` here; explicit close ordering is critical.

	var copyErr error // To capture error from io.Copy

	// Choose compression algorithm based on dst suffix.
	// Default to gzip if no recognized suffix.
	if strings.HasSuffix(dst, zstdSuffix) {
		enc, err := zstd.NewWriter(dstFile)
		if err != nil { // Error creating zstd writer
			_ = dstFile.Close()
			_ = os.Remove(tmpDst)
			return fmt.Errorf("failed to init zstd writer for %s: %v", dst, err)
		}
		_, copyErr = io.Copy(enc, srcFile) // Copy data from source file to zstd writer
		closeErr := enc.Close()            // Close zstd writer to flush data
		if copyErr == nil && closeErr != nil {
			copyErr = closeErr
		}
	} else {
		gz := gzip.NewWriter(dstFile)     // Default to gzip
		_, copyErr = io.Copy(gz, srcFile) // Copy data from source file to gzip writer
		closeErr := gz.Close()            // Close gzip writer to flush data
		if copyErr == nil && closeErr != nil {
			copyErr = closeErr
		}
	}

	if copyErr != nil { // Error during copy or close
		_ = dstFile.Close()
		_ = os.Remove(tmpDst)
		return fmt.Errorf("failed to write compressed data to %s: %w", dst, copyErr)
	}

	if err := dstFile.Close(); err != nil {
		_ = os.Remove(tmpDst)
		return fmt.Errorf("failed to close destination compressed file %s: %w", tmpDst, err)
	}

	// Atomically publish the completed compressed file.
	if err := l.resolvedRename(tmpDst, dst); err != nil {
		_ = os.Remove(tmpDst)
		return fmt.Errorf("failed to rename temp compressed file to %s: %w", dst, err)
	}

	if errChown := chown(dst, srcInfo); errChown != nil { // Attempt to chown the destination file
		// Log the chown error, but don't make it a fatal error for the compression process itself,
		// as the compressed file is valid. The original source file will still be removed.
		fmt.Fprintf(os.Stderr, "timberjack: [%s] failed to chown compressed log file %s: %v (source %s)\n",
			filepath.Base(src), dst, errChown, src)
		// Note: Depending on requirements, a chown failure could be considered critical.
		// For now, it's logged, and compression proceeds to remove the source.
	}

	// Windows file locking fix
	// Close the source file explicitly BEFORE attempting to remove it.
	// On Windows, you cannot delete an open file. The defer srcFile.Close() won't execute
	// until this function returns, so we must close it here before calling resolvedRemove().
	// This prevents "The process cannot access the file because it is being used" errors.
	if err := srcFile.Close(); err != nil {
		// Log the close error but continue with removal attempt
		fmt.Fprintf(os.Stderr, "timberjack: [%s] failed to close source file before removal: %v\n", filepath.Base(src), err)
	}

	// Finally, after successful compression and atomic rename, remove the original source file.
	if err = l.resolvedRemove(src); err != nil {
		// This is a more significant error if the original isn't removed, as it might be re-processed.
		return fmt.Errorf("failed to remove original source log file %s after compression: %w", src, err)
	}
	return nil // Compression successful
}

// effectiveCompression returns "none" | "gzip" | "zstd".
// Rule: if Compression is set, it wins; if empty, fallback to legacy Compress.
// Unknown strings default to "none" (and warn once).
func (l *Logger) effectiveCompression() string {
	alg := strings.ToLower(strings.TrimSpace(l.Compression))
	switch alg {
	case "gzip", "zstd":
		return alg
	case "none", "":
		if alg == "" && l.Compress {
			return "gzip"
		}
		return "none"
	default:
		fmt.Fprintf(os.Stderr, "timberjack: invalid compression %q — using none\n", alg)
		return "none"
	}
}

// compressedSuffix returns ".gz" / ".zst" or "" if none.
func (l *Logger) compressedSuffix() string {
	switch l.resolvedCompression {
	case "gzip":
		return compressSuffix
	case "zstd":
		return zstdSuffix
	default:
		return ""
	}
}

// trimCompressionSuffix strips one known compression suffix (".gz" or ".zst").
func trimCompressionSuffix(name string) string {
	name = strings.TrimSuffix(name, compressSuffix)
	name = strings.TrimSuffix(name, zstdSuffix)
	return name
}

// sanitizeReason turns an arbitrary string into a safe, short tag for filenames.
// Allowed: [a-z0-9_-]. Everything else becomes '-'. Collapses repeats, trims edges.
// Returns empty string if nothing usable remains.
func sanitizeReason(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" {
		return ""
	}
	const max = 32

	var b strings.Builder
	lastDash := false
	for _, r := range s {
		ok := (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_'
		if ok {
			b.WriteRune(r)
			lastDash = (r == '-')
		} else {
			// replace anything else (including whitespace) with a single '-'
			if !lastDash && b.Len() > 0 {
				b.WriteByte('-')
				lastDash = true
			}
		}
		if b.Len() >= max {
			break
		}
	}
	out := strings.Trim(b.String(), "-_")
	return out
}

// logInfo is a convenience struct to return the filename and its embedded
// timestamp, along with its os.FileInfo.
type logInfo struct {
	timestamp   time.Time // Parsed timestamp from the filename
	os.FileInfo           // Full FileInfo
}

// byFormatTime sorts a slice of logInfo structs by their parsed timestamp in descending order (newest first).
type byFormatTime []logInfo

func (b byFormatTime) Less(i, j int) bool {
	// Handle cases where timestamps might be zero (e.g., parsing failed, though timeFromName should error out)
	if b[i].timestamp.IsZero() && !b[j].timestamp.IsZero() {
		return false
	} // Treat zero time as oldest
	if !b[i].timestamp.IsZero() && b[j].timestamp.IsZero() {
		return true
	} // Non-zero is newer than zero
	if b[i].timestamp.IsZero() && b[j].timestamp.IsZero() {
		return false
	} // Equal if both are zero (order doesn't matter)
	return b[i].timestamp.After(b[j].timestamp) // Sort newest first
}
func (b byFormatTime) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byFormatTime) Len() int      { return len(b) }
