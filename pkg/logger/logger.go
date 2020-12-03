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

package logger // import "aws-observability.io/collector/logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var CollectorInstallPath = "/opt/aws/aws-otel-collector/"

var logfile = getLogFilePath()

var lumberjackLogger = &lumberjack.Logger{
	Filename:   logfile,
	MaxSize:    100, //MB
	MaxBackups: 5,   //backup files
	MaxAge:     7,   //days
	Compress:   true,
}

const (
	// WindowsLogFilePath represents the log file location on Windows.
	WindowsLogFilePath = "C:\\ProgramData\\Amazon\\AwsOTelCollector\\Logs\\"

	// UnixLogFilePath represents the log file location on Unix-y operating
	// systems like Linux and Mac.
	UnixLogFilePath = "/opt/aws/aws-otel-collector/logs/"
)

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

// getLogFilePath retuns the log file path depending on the OS.
func getLogFilePath() string {
	if runtime.GOOS == "windows" {
		CollectorInstallPath = WindowsLogFilePath
	}
	return CollectorInstallPath + "aws-otel-collector.log"
}

// SetLogLevel allow to set log level by environment vars
// possible values (DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL)
func SetLogLevel() {
	if level, ok := os.LookupEnv("AOT_LOG_LEVEL"); ok {
		os.Args = append(os.Args, "--log-level", level)
	}
}