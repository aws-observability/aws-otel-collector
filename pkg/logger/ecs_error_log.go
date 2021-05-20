package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"time"
)

type ECSErrorLogger struct {
	ErrorFilePath    string
	ErrorLogMaxAge   int
	ErrorFileMaxSize int

	queue     []zapcore.Entry
	ticker    time.Ticker
	errorFile *os.File
	size      int
}

// ECSErrorReporter func help select the different case between either write error log into file or
// rotate and delete old logs in file based on TTL
func (l *ECSErrorLogger) ECSErrorReporter() {
	defer l.ticker.Stop()
	for {
		select {
		// when new error add in the channel, this case will be selected and write it into error log file
		case e := <-errorLogChannel:
			l.Write(e)
		// this case will run every x mins (x is configurable) to delete the error which already expired, error log max age
		// depends on STATUS_MESSAGE_TTL
		case <-l.ticker.C:
			l.processTimeout()
		}
	}
}

func (l *ECSErrorLogger) Write(newError zapcore.Entry) {
	if l.errorFile == nil {
		if err := l.openExistingOrNewErrorFile(); err != nil {
			log.Printf("could not open error log file when write new error, err: %v", err)
		}
	}
	newErrorLog := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
		newError.Time, newError.Level, newError.Caller, newError.Message))

	if len(newErrorLog) > l.ErrorFileMaxSize {
		log.Printf("error size exceed the max size of error file")
		return
	}

	err := l.errorFile.Truncate(0)
	if err != nil {
		log.Printf("could not truncate error log file when write new error, err: %v", err)
	}
	l.size = 0

	errorNum := len(l.queue)
	// iterate errors in queue, if the old error's message same as the new error's message, skip it,
	// otherwise write it into the file and queue
	for i := 0; i < errorNum; i++ {
		e := l.queue[0]
		if e.Message != newError.Message {
			n, err := l.errorFile.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
				e.Time, e.Level, e.Caller, e.Message)))
			if err != nil {
				log.Printf("could not write error into error log file, err: %v", err)
			}
			l.size += n
			l.queue = append(l.queue, e)
		}
		l.queue = l.queue[1:]
	}
	// delete the old errors if old errors plus new error exceed the max size
	if l.size+len(newErrorLog) > l.ErrorFileMaxSize {
		l.rotate(newErrorLog)
	}

	l.queue = append(l.queue, newError)
	n, err := l.errorFile.Write(newErrorLog)
	if err != nil {
		log.Printf("could not write new error into error log file, err: %v", err)
	}
	l.size += n
}

// rotate func could delete the old logs to make sure error log file less than 1 KB
func (l *ECSErrorLogger) rotate(newErrorLog []byte) {
	file, err := os.OpenFile(l.ErrorFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("could not open error log file during rotation, err: %v", err)
	}

	err = file.Truncate(0)
	if err != nil {
		log.Printf("could not truncate error log file during rotation, err: %v", err)
	}

	errorNum := len(l.queue)
	for i := 0; i < errorNum; i++ {
		currentError := l.queue[0]
		currentErrorLog := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
			currentError.Time, currentError.Level, currentError.Caller, currentError.Message))
		if l.size+len(newErrorLog) > l.ErrorFileMaxSize {
			l.size -= len(currentErrorLog)
		} else {
			l.queue = append(l.queue, currentError)
			_, err = file.Write(currentErrorLog)
			if err != nil {
				log.Printf("could not write error into error log file during rotation, err: %v", err)
			}
		}
		l.queue = l.queue[1:]
	}
}

// processTimeout function could delete the expired error log from the file based on the ErrorLogMaxAge
func (l *ECSErrorLogger) processTimeout() {
	currentTime := time.Now()
	file, err := os.OpenFile(l.ErrorFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("could not open error log file during process timeout, err: %v", err)
	}

	err = file.Truncate(0)
	if err != nil {
		log.Printf("could not truncate error log file during rotation, err: %v", err)
	}
	l.size = 0

	errorNum := len(l.queue)
	for i := 0; i < errorNum; i++ {
		e := l.queue[0]
		if e.Time.Add(time.Duration(l.ErrorLogMaxAge) * time.Hour).After(currentTime) {
			n, err := file.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
				e.Time, e.Level, e.Caller, e.Message)))
			if err != nil {
				log.Printf("could not write error into error log file during process timeout, err: %v", err)
			}
			l.queue = append(l.queue, e)
			l.size += n
		}
		l.queue = l.queue[1:]
	}
}

// openExistingOrNewErrorFile opens the logfile if it exists. If there is no such file, a new file is created.
func (l *ECSErrorLogger) openExistingOrNewErrorFile() error {
	filename := l.ErrorFilePath
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return l.openNewErrorFile()
	}
	if err != nil {
		return fmt.Errorf("error getting error log file info: %s", err)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return l.openNewErrorFile()
	}
	l.errorFile = file
	return nil
}

func (l *ECSErrorLogger) openNewErrorFile() error {
	err := os.MkdirAll(l.errorDir(), 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new error logfile: %s", err)
	}
	name := l.ErrorFilePath
	_, err = os.Create(name)
	if err != nil {
		return fmt.Errorf("can't create new error logfile: %s", err)
	}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return fmt.Errorf("can't open new error logfile: %s", err)
	}
	l.errorFile = f
	return nil
}

func (l *ECSErrorLogger) errorDir() string {
	return filepath.Dir(l.ErrorFilePath)
}
