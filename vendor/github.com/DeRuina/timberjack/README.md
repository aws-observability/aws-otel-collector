# timberjack [![Go Reference](https://pkg.go.dev/badge/github.com/DeRuina/timberjack.svg)](https://pkg.go.dev/github.com/DeRuina/timberjack) [![Go Report Card](https://goreportcard.com/badge/github.com/DeRuina/timberjack)](https://goreportcard.com/report/github.com/DeRuina/timberjack) ![Audit](https://github.com/DeRuina/timberjack/actions/workflows/audit.yaml/badge.svg) ![Version](https://img.shields.io/github/v/tag/DeRuina/timberjack?sort=semver) [![Coverage Status](https://coveralls.io/repos/github/DeRuina/timberjack/badge.svg)](https://coveralls.io/github/DeRuina/timberjack) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)


### Timberjack is a Go package for writing logs to rolling files.

Timberjack is a forked and enhanced version of [`lumberjack`](https://github.com/natefinch/lumberjack), adding time-based rotation, clock-scheduled rotation, and opt-in compression (gzip or zstd).
Package `timberjack` provides a rolling logger with support for size-based and time-based log rotation.


## Installation

```bash
go get github.com/DeRuina/timberjack
```


## Import

```go
import "github.com/DeRuina/timberjack"
```

Timberjack is a pluggable component that manages log file writing and rotation. It works with any logger that writes to an `io.Writer`, including the standard library’s `log` package. `Logger` also satisfies `zapcore.WriteSyncer` (from `go.uber.org/zap`), so it can be passed directly to `zapcore.NewCore` or wrapped with `zapcore.AddSync`.

> ⚠️ Timberjack assumes **one process** writes to a given file. Reusing the same config from multiple
> processes on the same machine may lead to unexpected behavior.


## Example

To use timberjack with the standard library's `log` package, including interval-based and scheduled minute/daily rotation:

```go
import (
	"log"
	"time"
	"github.com/DeRuina/timberjack"
)

func main() {
	logger := &timberjack.Logger{
        Filename:          "/var/log/myapp/foo.log",    // Choose an appropriate path
        MaxSize:            500,                        // megabytes
        MaxBackups:         3,                          // backups
        MaxAge:             28,                         // days
        Compression:        "gzip",                     // "none" | "gzip" | "zstd" (preferred over legacy Compress)
        LocalTime:          true,                       // default: false (use UTC)
        RotationInterval:   24 * time.Hour,             // Rotate daily if no other rotation met
        RotateAtMinutes:    []int{0, 15, 30, 45},       // Also rotate at HH:00, HH:15, HH:30, HH:45
        RotateAt:           []string{"00:00", "12:00"}, // Also rotate at 00:00 and 12:00 each day
        BackupTimeFormat:   "2006-01-02-15-04-05",      // Rotated files will have format <logfilename>-2006-01-02-15-04-05-<reason>.log
        AppendTimeAfterExt: true,                       // put timestamp after ".log" (foo.log-<timestamp>-<reason>)
        FileMode:           0o644,                      // Custom permissions for newly created files. If unset or 0, defaults to 640.
	}
	log.SetOutput(logger)
	defer logger.Close() // Ensure logger is closed on application exit to stop goroutines

	log.Println("Application started")
	// ... your application logic ...
	log.Println("Application shutting down")
}
```

Manual rotation (e.g. on `SIGHUP`):

```go
import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/DeRuina/timberjack"
)


func main() {
    l := &timberjack.Logger{ Filename: "/var/log/myapp/foo.log" }
    log.SetOutput(l)
    defer l.Close()

    // Manual rotation on SIGHUP
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGHUP)

    go func() {
        for range c {
            // 1) Classic behavior: auto-pick "time" (if due) or "size"
            // _ = l.Rotate()

            // 2) New: tag the backup with your own reason
            _ = l.RotateWithReason("reload")
        }
    }()

    // ...
}
```



## Use with go.uber.org/zap

`Logger` satisfies `zapcore.WriteSyncer`, so you can plug it straight into a zap core:

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/DeRuina/timberjack"
)

func main() {
    w := &timberjack.Logger{
        Filename:   "/var/log/myapp/foo.log",
        MaxSize:    500,   // MB
        MaxBackups: 3,
        MaxAge:     28,    // days
        Compression: "zstd",
    }
    defer w.Close()

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
        zapcore.AddSync(w), // or pass w directly — Logger satisfies WriteSyncer
        zap.InfoLevel,
    )
    logger := zap.New(core)
    defer logger.Sync()

    logger.Info("Application started")
}
```


## Logger Configuration

```go
type Logger struct {
    Filename          string        // File to write logs to
    MaxSize           int           // Max size (MB) before rotation (default: 100)
    MaxAge            int           // Max age (days) to retain old logs
    MaxBackups        int           // Max number of backups to keep
    LocalTime         bool          // Use local time in rotated filenames

    // Compression controls post-rotation compression:
    //   "none" | "gzip" | "zstd"
    // Unknown/empty default to "none", unless the legacy Compress is true (see below).
    Compression       string

    // Deprecated: prefer Compression.
    // If Compression is empty and Compress is true = gzip compression.
    // Back-compat shim for old configs; will be removed in v2.
    Compress          bool


    RotationInterval  time.Duration // Rotate after this duration (if > 0)
    RotateAtMinutes   []int         // Specific minutes within an hour (0–59) to trigger rotation
    RotateAt          []string      // Specific daily times (HH:MM, 24-hour) to trigger rotation
    BackupTimeFormat  string        // Optional. If unset or invalid, defaults to 2006-01-02T15-04-05.000 (with fallback warning)
    AppendTimeAfterExt    bool      // if true, name backups like foo.log-<timestamp>-<reason> defaults to foo-<timestamp>-<reason>.log
    FileMode          os.FileMode   // Use custom permissions for newly created files. If zero (unset or 0), defaults to 640.
}
```


## How Rotation Works

1. **Size-Based**: If a write operation causes the current log file to exceed `MaxSize`, the file is rotated before the write. The backup filename will include `-size` as the reason.
2. **Time-Based (Interval)**: If `RotationInterval` is set (e.g., `24 * time.Hour` for daily rotation) and this duration has passed since the last rotation (of any type that updates the interval timer), the file is rotated upon the next write. The backup filename will include `-time` as the reason.
3. **Scheduled (Clock-Aligned)**: If `RotateAtMinutes` and/or `RotateAt` are configured (e.g., `[]int{0,30}` → rotate at `HH:00` and `HH:30`; or `[]string{"00:00"}` → rotate at midnight), a background goroutine triggers rotation at those times. These rotations use `-time` as the reason.
4. **Manual**: 
    - `Logger.Rotate()` forces rotation now. The backup reason will be `"time"` if an interval rotation is due, otherwise `"size"`.
    - `Logger.RotateWithReason("your-reason")` forces rotation and tags the backup with your **sanitized** reason (see below). If the provided reason is empty after sanitization, it falls back to the same behavior as Rotate().

Rotated files are renamed using the pattern:

By default, rotated files are named:
```
<name>-<timestamp>-<reason>.log
```

For example:

```
/var/log/myapp/foo-2025-04-30T15-00-00.000-size.log
/var/log/myapp/foo-2025-04-30T22-15-42.123-time.log
/var/log/myapp/foo-2025-05-01T10-30-00.000-time.log.gz  (if compressed)
```

If you prefer the extension to stay attached to the live name (better shell TAB completion),

set `AppendTimeAfterExt: true`:

```
<name>.log-<timestamp>-<reason>
```

For example:

```
/var/log/myapp/foo.log-2025-04-30T15-00-00.000-size
/var/log/myapp/foo.log-2025-04-30T22-15-42.123-time
/var/log/myapp/foo.log-2025-05-01T10-30-00.000-time.gz (if compressed)
```

Manual rotation with a custom reason `_ = logger.RotateWithReason("reload-now v2")`

For example:

```
foo-2025-05-01T10-30-00.000-reload-now-v2.log
```

### Compression

- Pick the algorithm with `Compression: "none" | "gzip" | "zstd"`.
- **Precedence**: If Compression is set, it **wins**. If it’s empty, legacy `Compress: true` means gzip; else no compression.
- Outputs use `.gz` or `.zst` suffix accordingly.
- Compression happens after rotation in a background goroutine.
- **Deprecation**: `Compress` is kept only for backward compatibility with old configs. It’s ignored when `Compression` is set. **It will be removed in v2**.

### Cleanup

On each new log file creation, timberjack:
- Deletes backups exceeding `MaxBackups` (keeps the newest rotations).
- Deletes backups older than `MaxAge` days.
- Compresses uncompressed backups if compression is enabled.

### Rotation modes at a glance

| Mode                           | Configure with                                | Trigger                                                             | Anchor                       | Background goroutine? | Rotates with zero writes? | Updates `lastRotationTime` | Backup suffix                                             | Notes                                                                                                             |
| ------------------------------ | --------------------------------------------- | ------------------------------------------------------------------- | ---------------------------- | :-------------------: | :-----------------------: | :------------------------: | --------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| **Size-based**                 | `MaxSize`                                     | A write would exceed `MaxSize`                                      | N/A                          |           No          |             No            |           **No**           | `-size`                                                   | Always active. A single write larger than `MaxSize` returns an error.                                             |
| **Interval-based**             | `RotationInterval > 0`                        | On **next write** after `now - lastRotationTime ≥ RotationInterval` | Duration since last rotation |           No          |             No            |     **Yes** (to `now`)     | `-time`                                                   | “Every N” rotations; not aligned to the wall clock.                                                               |
| **Scheduled minute-based**     | `RotateAtMinutes` (e.g. `[]int{0,30}`)        | At each `HH:MM` where minute matches                                | Clock minute marks           |        **Yes**        |          **Yes**          |           **Yes**          | `-time`                                                   | Expands minutes across all 24 hours. Invalid minutes are ignored **with a warning**. De-duplicated vs `RotateAt`. |
| **Scheduled daily fixed time** | `RotateAt` (e.g. `[]string{"00:00","12:00"}`) | At each listed `HH:MM` daily                                        | Clock minute marks           |        **Yes**        |          **Yes**          |           **Yes**          | `-time`                                                   | Ideal for “rotate at midnight”. De-duplicated vs `RotateAtMinutes`.                                               |
| **Manual**                     | `Logger.Rotate()`                             | When called                                                         | Immediate                    |           No          |            N/A            |           **No**           | `-time` if an interval rotation is due; otherwise `-size` | Handy for SIGHUP.
| **Manual (custom reason)**   | `Logger.RotateWithReason(s)`     | When called | Immediate | No | N/A | **No** | `-<sanitized reason>` | Falls back to `Rotate()` behavior if `s` sanitizes to empty. |

> **Time zone:** scheduling and filename timestamps use UTC by default, or local time if `LocalTime: true`.
> **Sanitized reason:** lowercase; `[a-z0-9_-]` only,  trims edge, max 32. 

## ⚠️ Rotation Notes & Warnings

* **`BackupTimeFormat` Values must be valid and should not change after initialization**  
  The `BackupTimeFormat` value **must be valid** and must follow the timestamp layout rules
  specified here: https://pkg.go.dev/time#pkg-constants. `BackupTimeFormat` supports more formats but it's recommended to use standard formats. If an **invalid** `BackupTimeFormat` is configured, Timberjack logs a warning to `os.Stderr` and falls back to the default format: `2006-01-02T15-04-05.000`. Rotation will still work, but the resulting filenames may not match your expectations.

* **Invalid `RotateAtMinutes`/`RotateAt` Values**  
  Values outside the valid range (`0–59`) for `RotateAtMinutes` or invalid time (`HH:MM`) for `RotateAt` or duplicates in `RotateAtMinutes`/`RotateAt` are ignored with a warning to stderr. Rotation continues with the valid schedule.

* **Logger Must Be Closed**
  Always call `logger.Close()` when done logging. This shuts down internal goroutines used for scheduled rotation and cleanup. Failing to close the logger can result in orphaned background processes, open file handles, and memory leaks.
  `Sync()` is safe to call at any point (including after `Close()`); it is a no-op when the logger is closed or no file has been opened yet.

* **Size-Based Rotation Is Always Active**
  Regardless of `RotationInterval` or `RotateAtMinutes`/`RotateAt`, size-based rotation is always enforced. If a write causes the log to exceed `MaxSize` (default: 100MB), it triggers an immediate rotation.

* **If Only `RotationInterval` Is Set**
  The logger rotates after the configured time has passed since the **last rotation**, regardless of file size.

* **If Only `RotateAtMinutes`/`RotateAt` Is Set**
  The logger rotates **at the clock times** specified, regardless of file size or duration passed. This is handled by a background goroutine. Rotated logs can be empty if no write has occurred.

* **If Both Are Set**  
  Both time-based strategies (`RotationInterval` and `RotateAtMinutes`) are evaluated. Whichever condition occurs first triggers rotation. However:

  * Both update the internal `lastRotationTime` field.
  * This means if a rotation happens due to `RotateAtMinutes`/`RotateAt`, it resets the interval timer, potentially **delaying or preventing** a `RotationInterval`-based rotation.

  This behavior ensures you won’t get redundant rotations, but it may make `RotationInterval` feel unpredictable if `RotateAtMinutes`/`RotateAt` is also configured.

## Contributing

We welcome contributions!
Please see our [contributing guidelines](CONTRIBUTING.md) before submitting a pull request.

## License

MIT
