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

package logger // import "aws-observability.io/collector/logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/aws-observability/aws-otel-collector/pkg/extraconfig"
)

var (
	UnixLogPath      = "/opt/aws/aws-otel-collector/logs/aws-otel-collector.log"
	WindowsLogPath   = "C:\\ProgramData\\Amazon\\AWSOTelCollector\\Logs\\aws-otel-collector.log"
	logfile          = getLogFilePath()
	lumberjackLogger = tryNewLumberJackLogger()
)

func tryNewLumberJackLogger() *lumberjack.Logger {
	if logfile != "" && !extraconfig.IsRunningInContainer() {
		return &lumberjack.Logger{
			Filename:   logfile,
			MaxSize:    100, //MB
			MaxBackups: 5,   //backup files
			MaxAge:     7,   //days
			Compress:   true,
		}
	}
	return nil
}

// WrapCoreOpt returns a zap.Option that wraps the provided core, teeing the output to the lumberjack writer.
// It uses a JSON encoder and the same level as the provided core.
// If the lumberjack logger is not configured returns the provided core unmodified.
func WrapCoreOpt() zap.Option {
	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		if lumberjackLogger == nil {
			return core
		}

		encoderConfig := zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			CallerKey:      "caller",
			StacktraceKey:  "stack",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		return zapcore.NewTee(core, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(lumberjackLogger), core.(zapcore.LevelEnabler)))
	})
}

// SetupErrorLogger setup lumberjackLogger for go logger
func SetupErrorLogger() {
	var writer io.WriteCloser
	// When running in container, always log to stderr, it makes debugging easier.
	if lumberjackLogger != nil {
		err := os.MkdirAll(filepath.Dir(logfile), 0750)
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
		return WindowsLogPath
	}
	return UnixLogPath
}

// SetLogLevel allows to set log level by parameter
// possible values (DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL)
func SetLogLevel(level string) {
	if level != "" {
		os.Args = append(os.Args, fmt.Sprintf("--config=yaml:service::telemetry::logs::level: %s", level))
	}
}
