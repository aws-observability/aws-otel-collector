/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"aws-observability.io/collector/pkg/consts"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logfile = getLogFilePath()

var lumberjackLogger = &lumberjack.Logger{
	Filename:   logfile,
	MaxSize:    100, //MB
	MaxBackups: 5,   //backup files
	MaxAge:     7,   //days
	Compress:   true,
}

// GetLumberHook returns lumberjackLogger as a Zap hook
// for processing log size and log rotation
func GetLumberHook() func(e zapcore.Entry) error {
	return func(e zapcore.Entry) error {
		_, err := lumberjackLogger.Write(
			[]byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v, Stack:%+v}\r\n",
				e.Time, e.Level, e.Caller, e.Message, e.Stack)))
		if err != nil {
			return err
		}
		return nil
	}
}

// SetupErrorLogger setup lumberjackLogger for go logger
func SetupErrorLogger() {
	log.SetFlags(0)
	var writer io.WriteCloser
	if logfile != "" {
		err := os.MkdirAll(filepath.Dir(logfile), 0755)
		if err != nil {
			log.Printf("D! fail to chmod on log file due to : %v \n", err)
		}
		// The codes below should not change, because the retention information has already been published to public doc.
		writer = lumberjackLogger
	} else {
		writer = os.Stderr
	}
	log.SetOutput(writer)
}

//this method retuns the log file path depending on the OS
func getLogFilePath() string {
	if runtime.GOOS == "windows" {
		return consts.WIN_LOGFILE_PATH
	} else {
		return consts.LINUX_LOGFILE_PATH
	}
}
