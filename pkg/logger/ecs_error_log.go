/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ECSErrorLogger struct {
	ErrorLogFilePath    string
	ErrorLogMaxAge      int
	ErrorLogFileMaxSize int

	errorChannel chan zapcore.Entry
	queue        []zapcore.Entry
	ticker       time.Ticker
	errorFile    *os.File
	size         int
}

func NewECSErrorLogger() *ECSErrorLogger {
	ecsErrorLogFilePath := os.Getenv("ECS_ERROR_LOG_FILE_PATH")
	errorLogMaxAge, err := strconv.Atoi(os.Getenv("ECS_ERROR_LOG_TTL"))
	if err != nil {
		log.Fatalf("failed to get STATUS_MESSAGE_TTL: %v", err)
	}
	errorFileSize, err := strconv.Atoi(os.Getenv("ECS_ERROR_LOG_MAX_BYTE_LENGTH"))
	if err != nil {
		log.Fatalf("failed to get STATUS_MESSAGE_MAX_BYTE_LENGTH: %v", err)
	}
	return &ECSErrorLogger{
		ErrorLogFilePath:    ecsErrorLogFilePath,
		ErrorLogMaxAge:      errorLogMaxAge,
		ErrorLogFileMaxSize: errorFileSize,
		errorChannel:        make(chan zapcore.Entry, 100),
	}
}

// Run func help select the different case between either write error log into file or
// rotate and delete old logs in file based on TTL
func (l *ECSErrorLogger) Run() {
	defer l.ticker.Stop()
	for {
		select {
		case newError := <-l.errorChannel:
			l.Write(newError)
		case <-l.ticker.C:
			l.processTimeout()
		}
	}
}

func (l *ECSErrorLogger) Write(newError zapcore.Entry) {
	if l.errorFile == nil {
		if err := l.openExistingOrNewErrorFile(); err != nil {
			log.Printf("[ecs error reporter] could not open error log file when write new error, err: %v", err)
			return
		}
	}
	newErrorLog := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
		newError.Time, newError.Level, newError.Caller, newError.Message))

	if len(newErrorLog) > l.ErrorLogFileMaxSize {
		log.Printf("[ecs error reporter] error size exceed the max size of error file")
		return
	}

	err := l.errorFile.Truncate(0)
	if err != nil {
		log.Printf("[ecs error reporter] could not truncate error log file when write new error, err: %v", err)
		return
	}
	l.size = 0

	errorNum := len(l.queue)
	for i := 0; i < errorNum; i++ {
		e := l.queue[0]
		if e.Message != newError.Message {
			n, err := l.errorFile.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
				e.Time, e.Level, e.Caller, e.Message)))
			if err != nil {
				log.Printf("[ecs error reporter] could not write error into error log file, err: %v", err)
				return
			}
			l.size += n
			l.queue = append(l.queue, e)
		}
		l.queue = l.queue[1:]
	}

	if l.size+len(newErrorLog) > l.ErrorLogFileMaxSize {
		l.rotate(len(newErrorLog))
	}

	l.queue = append(l.queue, newError)
	n, err := l.errorFile.Write(newErrorLog)
	if err != nil {
		log.Printf("[ecs error reporter] could not write new error into error log file, err: %v", err)
		return
	}
	l.size += n
}

// rotate func could delete the old logs to make sure error log file less than 1 KB
func (l *ECSErrorLogger) rotate(newErrorLogSize int) {
	file, err := os.OpenFile(l.ErrorLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("[ecs error reporter] could not open error log file during rotation, err: %v", err)
		return
	}

	err = file.Truncate(0)
	if err != nil {
		log.Printf("[ecs error reporter] could not truncate error log file during rotation, err: %v", err)
		return
	}

	errorNum := len(l.queue)
	for i := 0; i < errorNum; i++ {
		currentError := l.queue[0]
		currentErrorLog := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
			currentError.Time, currentError.Level, currentError.Caller, currentError.Message))
		if l.size+newErrorLogSize > l.ErrorLogFileMaxSize {
			l.size -= len(currentErrorLog)
		} else {
			l.queue = append(l.queue, currentError)
			_, err = file.Write(currentErrorLog)
			if err != nil {
				log.Printf("[ecs error reporter] could not write error into error log file during rotation, err: %v", err)
				return
			}
		}
		l.queue = l.queue[1:]
	}
}

// processTimeout function could delete the expired error log from the file based on the ErrorLogMaxAge
func (l *ECSErrorLogger) processTimeout() {
	currentTime := time.Now()
	file, err := os.OpenFile(l.ErrorLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("[ecs error reporter] could not open error log file during process timeout, err: %v", err)
		return
	}

	err = file.Truncate(0)
	if err != nil {
		log.Printf("[ecs error reporter] could not truncate error log file during rotation, err: %v", err)
		return
	}
	l.size = 0

	errorNum := len(l.queue)
	for i := 0; i < errorNum; i++ {
		e := l.queue[0]
		if e.Time.Add(time.Duration(l.ErrorLogMaxAge) * time.Hour).After(currentTime) {
			n, err := file.Write([]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
				e.Time, e.Level, e.Caller, e.Message)))
			if err != nil {
				log.Printf("[ecs error reporter] could not write error into error log file during process timeout, err: %v", err)
				return
			}
			l.queue = append(l.queue, e)
			l.size += n
		}
		l.queue = l.queue[1:]
	}
}

// GetErrorHook function could generate a zap hook
func (l *ECSErrorLogger) GetErrorHook() func(e zapcore.Entry) error {
	return func(e zapcore.Entry) error {
		if e.Level >= zapcore.ErrorLevel && os.Getenv("STATUS_MESSAGE_FILE_PATH") != "" {
			l.errorChannel <- e
		}
		return nil
	}
}

// openExistingOrNewErrorFile opens the logfile if it exists. If there is no such file, a new file is created.
func (l *ECSErrorLogger) openExistingOrNewErrorFile() error {
	filename := l.ErrorLogFilePath
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
	name := l.ErrorLogFilePath
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
	return filepath.Dir(l.ErrorLogFilePath)
}
