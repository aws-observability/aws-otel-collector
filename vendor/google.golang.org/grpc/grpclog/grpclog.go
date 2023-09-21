/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package grpclog defines logging for grpc.
//
// All logs in transport and grpclb packages only go to verbose level 2.
// All logs in other packages in grpc are logged in spite of the verbosity level.
//
// In the default logger,
// severity level can be set by environment variable GRPC_GO_LOG_SEVERITY_LEVEL,
// verbosity level can be set by GRPC_GO_LOG_VERBOSITY_LEVEL.
package grpclog // import "google.golang.org/grpc/grpclog"

import (
	"os"

	"google.golang.org/grpc/internal/grpclog"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return grpclog.Logger.V(l)
}

// Info logs to the INFO log.
<<<<<<< HEAD
func Info(args ...interface{}) {
=======
func Info(args ...any) {
>>>>>>> main
	grpclog.Logger.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
<<<<<<< HEAD
func Infof(format string, args ...interface{}) {
=======
func Infof(format string, args ...any) {
>>>>>>> main
	grpclog.Logger.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
<<<<<<< HEAD
func Infoln(args ...interface{}) {
=======
func Infoln(args ...any) {
>>>>>>> main
	grpclog.Logger.Infoln(args...)
}

// Warning logs to the WARNING log.
<<<<<<< HEAD
func Warning(args ...interface{}) {
=======
func Warning(args ...any) {
>>>>>>> main
	grpclog.Logger.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
<<<<<<< HEAD
func Warningf(format string, args ...interface{}) {
=======
func Warningf(format string, args ...any) {
>>>>>>> main
	grpclog.Logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
<<<<<<< HEAD
func Warningln(args ...interface{}) {
=======
func Warningln(args ...any) {
>>>>>>> main
	grpclog.Logger.Warningln(args...)
}

// Error logs to the ERROR log.
<<<<<<< HEAD
func Error(args ...interface{}) {
=======
func Error(args ...any) {
>>>>>>> main
	grpclog.Logger.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
<<<<<<< HEAD
func Errorf(format string, args ...interface{}) {
=======
func Errorf(format string, args ...any) {
>>>>>>> main
	grpclog.Logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
<<<<<<< HEAD
func Errorln(args ...interface{}) {
=======
func Errorln(args ...any) {
>>>>>>> main
	grpclog.Logger.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
<<<<<<< HEAD
func Fatal(args ...interface{}) {
=======
func Fatal(args ...any) {
>>>>>>> main
	grpclog.Logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
<<<<<<< HEAD
func Fatalf(format string, args ...interface{}) {
=======
func Fatalf(format string, args ...any) {
>>>>>>> main
	grpclog.Logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
<<<<<<< HEAD
func Fatalln(args ...interface{}) {
=======
func Fatalln(args ...any) {
>>>>>>> main
	grpclog.Logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
<<<<<<< HEAD
func Print(args ...interface{}) {
=======
func Print(args ...any) {
>>>>>>> main
	grpclog.Logger.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
<<<<<<< HEAD
func Printf(format string, args ...interface{}) {
=======
func Printf(format string, args ...any) {
>>>>>>> main
	grpclog.Logger.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
<<<<<<< HEAD
func Println(args ...interface{}) {
=======
func Println(args ...any) {
>>>>>>> main
	grpclog.Logger.Infoln(args...)
}
