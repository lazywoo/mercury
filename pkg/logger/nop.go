package logger

type NopLogger struct{}

func (n *NopLogger) Info(msg string, args ...Field) {}

func (n *NopLogger) Debug(msg string, args ...Field) {}

func (n *NopLogger) Warn(msg string, args ...Field) {}

func (n *NopLogger) Error(msg string, args ...Field) {}

func NewNopLogger() Logger {
	return &NopLogger{}
}
