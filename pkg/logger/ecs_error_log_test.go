package logger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testErrorLogFilePath = "/Users/tianduo/Documents/error-reporting-test.log"

func TestWriteError(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	openExistingOrNewErrorFile = OpenExistOrNew
	entry := zapcore.Entry{
		Message: "test",
	}
	errorFileSize = 1024
	_, err := WriteError(entry, testErrorLogFilePath)
	queue = queue[:len(queue)-1]
	require.NoError(t, err)
	errorLog, err := ioutil.ReadFile(testErrorLogFilePath)
	if err != nil {
		return
	}
	s := string(errorLog)
	assert.True(t, strings.Contains(s, "test"))
}

func TestWriteSecondErrorIfSizeExceed(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	openExistingOrNewErrorFile = OpenExistOrNew
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
	errorFileSize = 1024
	_, err := WriteError(e1, testErrorLogFilePath)
	_, err = WriteError(e2, testErrorLogFilePath)
	require.NoError(t, err)
	queue = queue[:len(queue)-1]
	errorLog, err := ioutil.ReadFile(testErrorLogFilePath)
	if err != nil {
		return
	}
	s := string(errorLog)
	assert.True(t, strings.Contains(s, "a"))
	assert.True(t, !strings.Contains(s, "test"))
}

func TestRotate(t *testing.T) {
	defer os.Remove(testErrorLogFilePath)
	testErrorLogFile, _ := os.OpenFile(testErrorLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	errormessage1 := make([]byte, 700)
	errormessage2 := make([]byte, 700)
	for i := 0; i < 700; i++ {
		errormessage1[i] = 'a'
		errormessage2[i] = 'b'
	}
	e1 := zapcore.Entry{
		Message: string(errormessage1),
	}
	e2 := zapcore.Entry{
		Message: string(errormessage2),
	}
	content1 := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
		e1.Time, e1.Level, e1.Caller, e1.Message))
	content2 := []byte(fmt.Sprintf("{%+v, Level:%+v, Caller:%+v, Message:%+v}\r\n",
		e2.Time, e2.Level, e2.Caller, e2.Message))
	_, _ = testErrorLogFile.Write(content1)
	_, _ = testErrorLogFile.Write(content2)
	queue = append(queue, e1)
	queue = append(queue, e2)
	errorFileSize = 1024
	byteNum := len(content1) + len(content2)
	n, err := rotate(byteNum, testErrorLogFilePath)
	queue = queue[:len(queue)-1]
	require.NoError(t, err)
	assert.Equal(t, len(content2), n)
}

func OpenExistOrNew() error {
	testErrorFilePath := filepath.Dir(testErrorLogFilePath)
	err := os.MkdirAll(testErrorFilePath, 0744)
	if err != nil {
		return err
	}
	_, err = os.Create(testErrorLogFilePath)
	if err != nil {
		return err
	}
	errorFile, err = os.OpenFile(testErrorLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	return nil
}
