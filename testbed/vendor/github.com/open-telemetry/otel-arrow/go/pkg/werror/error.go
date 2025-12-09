/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package werror

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Wrapper wraps an error with the file, line, and function where the error
// was wrapped and an optional context map.
type Wrapper struct {
	// The wrapped error.
	err error

	// The file, line, and function where the error was wrapped.
	file     string
	line     int
	function string

	// An optional context map.
	context map[string]interface{}
}

// Error returns the wrapped error's message.
func (w Wrapper) Error() string {
	var msg strings.Builder

	msg.WriteString(w.function)
	msg.WriteString(":")
	msg.WriteString(strconv.Itoa(w.line))

	if w.context != nil {
		msg.WriteString("{")
		i := 0
		for k, v := range w.context {
			if i > 0 {
				msg.WriteString(", ")
			}
			msg.WriteString(k)
			msg.WriteString("=")
			msg.WriteString(fmt.Sprintf("%v", v))
			i++
		}
		msg.WriteString("}")
	}

	if w.err != nil {
		msg.WriteString("->")
		msg.WriteString(w.err.Error())
	}

	return msg.String()
}

// Unwrap returns the wrapped error.
func (w Wrapper) Unwrap() error {
	return w.err
}

// File returns the file where the error was wrapped.
func (w Wrapper) File() string {
	return w.file
}

// Line returns the line where the error was wrapped.
func (w Wrapper) Line() int {
	return w.line
}

// Function returns the function where the error was wrapped.
func (w Wrapper) Function() string {
	return w.function
}

// Wrap wraps the given error with the current file, line, and function.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return Wrapper{
		err:      err,
		file:     file,
		line:     line,
		function: fn.Name(),
		context:  nil,
	}
}

// WrapWithContext wraps the given error with the current file, line, function,
// and the given context.
func WrapWithContext(err error, context map[string]interface{}) error {
	if err == nil {
		return nil
	}

	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return Wrapper{
		err:      err,
		file:     file,
		line:     line,
		function: fn.Name(),
		context:  context,
	}
}

func WrapWithMsg(err error, msg string) error {
	if err == nil {
		return nil
	}

	context := map[string]interface{}{"msg": msg}

	if err == nil {
		return nil
	}

	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return Wrapper{
		err:      err,
		file:     file,
		line:     line,
		function: fn.Name(),
		context:  context,
	}
}
