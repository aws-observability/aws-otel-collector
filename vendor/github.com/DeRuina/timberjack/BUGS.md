# Bug Tracker

Tracks known bugs, recently fixed bugs, and open issues found during audits.

---

## Fixed (unreleased — pending next release)

These were found and fixed in a pre-release audit session on 2026-05-26.
All fixes are in `timberjack.go` with tests in `timberjack_test.go`.

### 1. Write-after-close ignores configured `FileMode` and bypasses `osOpenFile`

**Location:** `Write()` closed-logger path (~line 288)
**Severity:** High (security — file created with wrong permissions)
**Fix:** Use `l.resolvedOpenFile` and derive mode from `l.FileMode` / `defaultFileMode`.

### 2. Write-after-close has no `MaxSize` guard

**Location:** `Write()` closed-logger path (~line 284)
**Severity:** Medium (behavioral — oversized payloads silently accepted)
**Fix:** Added the same `writeLen > l.max()` check as the normal write path.

### 3. Scheduled rotation goroutine sets `lastRotationTime` to wall time instead of the mark

**Location:** `runScheduledRotations()` (~line 554)
**Severity:** Medium (correctness — can cause duplicate rotations with mocked clocks)
**Fix:** Changed `l.lastRotationTime = nowFn()` to `l.lastRotationTime = nextRotationAbsoluteTime`.

### 4. `sanitizeReason` produces `_-` sequences when `_` is followed by a space

**Location:** `sanitizeReason()` (~line 1371)
**Severity:** Low (cosmetic — garbled backup filenames like `my_-reason.log`)
**Fix:** Changed `lastDash = (r == '-')` to `lastDash = (r == '-' || r == '_')`.

### 5. `openExistingOrNew` calls `os.OpenFile` directly instead of `osOpenFile` global

**Location:** `openExistingOrNew()` (~line 862)
**Severity:** Low (test isolation — append-open path bypasses mock-able global)
**Fix:** Changed `os.OpenFile(...)` to `osOpenFile(...)`.

---

## Open — Found in pre-release audit (2026-05-26)

### OPEN-1: Backup files with hyphens in the reason are never cleaned up ⚠️

**Location:** `timeFromName()` (~line 1111)
**Severity:** High
**Root cause:** `strings.LastIndex(trimmed, "-")` splits on the *last* hyphen to separate timestamp from reason. `sanitizeReason` permits `-` in output, so `RotateWithReason("my-feature")` produces `app-...-my-feature.log`. `LastIndex` finds the `-` before `feature`, making `ts = "...-my"`, which fails `time.ParseInLocation`. The file is silently invisible to `oldLogFiles()` — never counted against `MaxBackups` or `MaxAge` — and accumulates indefinitely.

**Reproducer:** Call `RotateWithReason("my-tag")` with `MaxBackups: 3`. The resulting backup is never removed regardless of how many rotations occur.

**Proposed fix:** Use a delimiter that cannot appear in a valid timestamp and is stripped from reasons — e.g. double-underscore `__` between timestamp and reason. Requires a migration path for existing backup files.

---

### OPEN-2: Re-compression after failed source removal clobbers valid compressed backup

**Location:** `compressLogFile()` (~line 1318)
**Severity:** Medium
**Root cause:** After the atomic rename `tmpDst → dst` succeeds, `resolvedRemove(src)` can fail (e.g. `EACCES`). `compressLogFile` returns an error. On the next mill cycle `src` still exists without a `.gz`/`.zst` suffix, so it is added to `filesToCompress` again — overwriting the already-valid `dst` with a fresh compression of `src`.

**Proposed fix:** Before compressing, check whether `dst` already exists; if so, only attempt to remove `src` and skip recompression.

---

### OPEN-3: `l.Filename` read without `l.mu` in scheduled-rotation goroutine

**Location:** `runScheduledRotations()` (~line 540)
**Severity:** Low (data race under `-race` if caller mutates `Filename` after first write)
**Root cause:** The `!foundNextSlot` error branch reads `l.Filename` with no lock held inside the background goroutine. If a caller changes `l.Filename` concurrently the race detector fires.

**Proposed fix:** Snapshot `l.Filename` at goroutine start time alongside the existing `slots`/`loc`/`nowFn` snapshots and use it in all `fmt.Fprintf` calls inside the goroutine.

---

### OPEN-4: Partial rotation failure leaves `l.logStartTime` inconsistent with backup file timestamp

**Location:** `rotate()` / `openNew()` (~line 663)
**Severity:** Low (timing skew, no data loss)
**Root cause:** If `openNew()` successfully renames the current log to a backup (stamping it with `rotationTimeForBackup`) but then fails to create the new log file (e.g. disk full), `rotate()` returns an error with `l.file == nil`. The next `Write → openExistingOrNew → openNew("initial")` hits `os.IsNotExist` and sets `l.logStartTime = now` — inconsistent with the backup already on disk. This can skew interval-rotation timing but does not lose data.

**Proposed fix:** On `openNew` create-file failure, attempt to rename the backup back to the original name before returning the error to restore a consistent state.
