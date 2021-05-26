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
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

var testErrorLogFilePath = "/opt/aws/aws-otel-collector/logs/ecs-error-logger-test.log"

func TestWrite(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	entry := zapcore.Entry{
		Message: "test",
	}
	l := &ECSErrorLogger{
		ErrorFilePath:    testErrorLogFilePath,
		ErrorLogMaxAge:   10,
		ErrorFileMaxSize: 1024,
	}
	l.Write(entry)
	errorLog, err := ioutil.ReadFile(testErrorLogFilePath)
	if err != nil {
		return
	}
	s := string(errorLog)
	assert.True(t, strings.Contains(s, "test"))
}

func TestWriteSecondErrorIfSizeExceed(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	l := &ECSErrorLogger{
		ErrorFilePath:    testErrorLogFilePath,
		ErrorLogMaxAge:   10,
		ErrorFileMaxSize: 1024,
	}
	e1 := zapcore.Entry{
		Message: "test",
	}
	errormessage := make([]byte, 900)
	for i := 0; i < 900; i++ {
		errormessage[i] = 'a'
	}
	e2 := zapcore.Entry{
		Message: string(errormessage),
	}
	l.Write(e1)
	l.Write(e2)
	errorLog, err := ioutil.ReadFile(testErrorLogFilePath)
	if err != nil {
		return
	}
	s := string(errorLog)
	assert.True(t, strings.Contains(s, "aaa"))
	assert.True(t, !strings.Contains(s, "test"))
}

func TestRotate(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	l := &ECSErrorLogger{
		ErrorFilePath:    testErrorLogFilePath,
		ErrorLogMaxAge:   10,
		ErrorFileMaxSize: 1024,
	}
	error1, error2, error3 := make([]byte, 400), make([]byte, 400), make([]byte, 400)
	for i := 0; i < 400; i++ {
		error1[i], error2[i], error3[i] = 'a', 'b', 'c'
	}
	e1 := zapcore.Entry{
		Message: string(error1),
	}
	e2 := zapcore.Entry{
		Message: string(error2),
	}
	e3 := zapcore.Entry{
		Message: string(error3),
	}
	content3 := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
		e3.Time, e3.Level, e3.Caller, e3.Message))
	l.Write(e1)
	l.Write(e2)
	l.rotate(len(content3))
	errorLog, err := ioutil.ReadFile(testErrorLogFilePath)
	if err != nil {
		return
	}
	s := string(errorLog)
	assert.True(t, !strings.Contains(s, "aaa"))
	assert.True(t, strings.Contains(s, "bbb"))
}

var fakeCurrentTime = time.Now()

func fakeTime() time.Time {
	return fakeCurrentTime
}

func TestProcessTimeOut(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	l := &ECSErrorLogger{
		ErrorFilePath:    testErrorLogFilePath,
		ErrorLogMaxAge:   10,
		ErrorFileMaxSize: 1024,
	}
	time1 := fakeTime().Add(-20 * time.Hour)
	time2 := fakeTime().Add(-15 * time.Hour)
	time3 := fakeTime().Add(-5 * time.Hour)
	e1 := zapcore.Entry{
		Time:    time1,
		Message: "test1",
	}
	e2 := zapcore.Entry{
		Time:    time2,
		Message: "test2",
	}
	e3 := zapcore.Entry{
		Time:    time3,
		Message: "test3",
	}
	l.Write(e1)
	l.Write(e2)
	l.Write(e3)
	l.processTimeout()
	errorLog, err := ioutil.ReadFile(testErrorLogFilePath)
	if err != nil {
		return
	}
	s := string(errorLog)
	assert.True(t, !strings.Contains(s, "test1"))
	assert.True(t, !strings.Contains(s, "test2"))
	assert.True(t, strings.Contains(s, "test3"))
}
