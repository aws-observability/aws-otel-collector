package log

// Fields is a map that is used to populated logging context.
type Fields map[string]interface{}

type nilLogger struct {
}

func (n nilLogger) Debug(msg string) {
}

func (n nilLogger) Warn(msg string) {
}

func (n nilLogger) Error(msg string) {
}

func (n nilLogger) Info(msg string) {
}

func (n nilLogger) Panic(msg string) {
}

func (n nilLogger) WithFields(fields Fields) Logger {
	return nilLogger{}
}

func (n nilLogger) WithError(err error) Logger {
	return nilLogger{}
}

// Nil logger is a silent logger interface.
var Nil = nilLogger{}

var _ Logger = (*nilLogger)(nil)

// Logger is generic logging interface.
type Logger interface {
	Debug(msg string)
	Warn(msg string)
	Error(msg string)
	Info(msg string)
	Panic(msg string)
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}
