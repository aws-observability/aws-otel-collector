package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var ErrorFilePath = os.Getenv("STATUS_MESSAGE_FILE_PATH")

var errorLogMaxAge = os.Getenv("STATUS_MESSAGE_TTL")
var errorLogMaxAgeInHours, _ = strconv.Atoi(errorLogMaxAge)

var size = os.Getenv("STATUS_MESSAGE_MAX_BYTE_LENGTH")
var errorFileSize, _ = strconv.Atoi(size)

var queue []zapcore.Entry
var errorLogChannel = make(chan zapcore.Entry, 10)
var ticker = time.NewTicker(10 * time.Minute)
var errorFile *os.File

// ECSErrorReporter func help select the different case between either write error log into file or
// rotate and delete old logs in file based on TTL
func ECSErrorReporter() {
	defer ticker.Stop()
	for {
		select {
		// when new error add in the channel, this case will write it into file
		case e := <-errorLogChannel:
			_, _ = WriteError(e, ErrorFilePath)
		// this case will run every x mins (x is configurable) to delete the error which already expired, error log max age
		// depends on STATUS_MESSAGE_TTL
		case <-ticker.C:
			currentTime := time.Now()
			file, _ := os.OpenFile(ErrorFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
			_ = file.Truncate(0)
			errorNum := len(queue)
			for i := 0; i < errorNum; i++ {
				e := queue[0]
				if e.Time.Add(time.Duration(errorLogMaxAgeInHours) * time.Hour).After(currentTime) {
					_, _ = file.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
						e.Time, e.Level, e.Caller, e.Message)))
				} else {
					queue = queue[1:]
				}
			}
			_ = file.Close()
		}
	}
}

// WriteError func could write the new error log message into file
func WriteError(newErrorLog zapcore.Entry, errorLogfile string) (n int, err error) {
	var byteNum int
	if errorFile == nil {
		if err = openExistingOrNewErrorFile(); err != nil {
			return 0, err
		}
	}
	err = errorFile.Truncate(0)
	if err != nil {
		return byteNum, err
	}
	// iterate errors in queue, if the old error's message same as the new error's message, skip it,
	// otherwise write it into the file and queue
	var errorNum = len(queue)
	for i := 0; i < errorNum; i++ {
		e := queue[0]
		if e.Message != newErrorLog.Message {
			n, err = errorFile.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
				e.Time, e.Level, e.Caller, e.Message)))
			if err != nil {
				return byteNum, err
			}
			byteNum += n
			queue = append(queue, e)
		}
		queue = queue[1:]
	}
	queue = append(queue, newErrorLog)
	n, err = errorFile.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
		newErrorLog.Time, newErrorLog.Level, newErrorLog.Caller, newErrorLog.Message)))
	if err != nil {
		return byteNum, err
	}
	byteNum += n
	// errorFileSize should keep 1 KB
	if byteNum >= errorFileSize {
		byteNum, err = rotate(byteNum, errorLogfile)
		if err != nil {
			return byteNum, err
		}
	}
	return byteNum, err
}

// rotate func could delete the old logs to make sure error log file less than 1 KB
func rotate(byteNum int, errorLogfile string) (n int, err error) {
	file, err := os.OpenFile(errorLogfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return byteNum, err
	}
	err = file.Truncate(0)
	if err != nil {
		return byteNum, err
	}
	errorNum := len(queue)
	for i := 0; i < errorNum; i++ {
		currentError := queue[0]
		currentLog := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
			currentError.Time, currentError.Level, currentError.Caller, currentError.Message))
		if byteNum > errorFileSize {
			byteNum -= len(currentLog)
			queue = queue[1:]
		} else {
			_, err = file.Write(currentLog)
			if err != nil {
				return byteNum, err
			}
		}
	}
	return byteNum, err
}

// openExistingOrNewErrorFile opens the logfile if it exists. If there is no such file, a new file is created.
var openExistingOrNewErrorFile = func() error {
	filename := ErrorFilePath
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return openNewErrorFile()
	}
	if err != nil {
		return fmt.Errorf("error getting error log file info: %s", err)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return openNewErrorFile()
	}
	errorFile = file
	return nil
}

func openNewErrorFile() error {
	err := os.MkdirAll(errorDir(), 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new error logfile: %s", err)
	}
	name := errorFileName()
	_, err = os.Create(name)
	if err != nil {
		return fmt.Errorf("can't create new error logfile: %s", err)
	}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return fmt.Errorf("can't open new error logfile: %s", err)
	}
	errorFile = f
	return nil
}

func errorDir() string {
	return filepath.Dir(ErrorFilePath)
}

func errorFileName() string {
	if ErrorFilePath != "" {
		return ErrorFilePath
	}
	name := filepath.Base(os.Args[0]) + "-lumberjack.log"
	return filepath.Join(os.TempDir(), name)
}
