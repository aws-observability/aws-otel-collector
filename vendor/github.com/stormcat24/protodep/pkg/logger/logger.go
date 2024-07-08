package logger

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Info(format string, a ...interface{}) {
	color.Green("[INFO] "+format, a...)
}

func Warn(format string, a ...interface{}) {
	color.Yellow("[WARN] "+format, a...)
}

func Error(format string, a ...interface{}) {
	color.Red("[ERROR] "+format, a...)
}

type spinnerWrapper struct {
	spinner *spinner.Spinner
}

func (s *spinnerWrapper) Stop() {
	if s.spinner != nil {
		s.spinner.Stop()
	}
}

func (s *spinnerWrapper) Finish() {
	if s.spinner != nil {
		s.spinner.Stop()
	}
	fmt.Print("\n")
}

func InfoWithSpinner(format string, a ...interface{}) *spinnerWrapper {
	txt := color.GreenString("[INFO] "+format, a...)
	fmt.Print(txt)

	var s *spinner.Spinner
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Print("\n")
		s = spinner.New(spinner.CharSets[38], 100*time.Millisecond) // Build our new spinner
		s.Start()
	}

	return &spinnerWrapper{s}
}
