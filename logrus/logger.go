// Package logrus implements log.ContextLogger using logrus logger as a
// back-end.
//
// Example usage:
//
// import (
//     "github.com/go-mixins/log"
//     "github.com/go-mixins/log/logrus"
// )
// ...
// type App struct {
//     log.Logger
// }
// ...
// logger := logrus.New()
// logger.Level = logrus.DebugLevel
// ...
// app := &App{
//      Logger: logger,
// }
// ...
// app.Info("server started")
//
package logrus

import (
	"github.com/go-mixins/log"
	logrus "github.com/sirupsen/logrus"
)

// Some useful constants from the Logrus to reduce import clutter
const (
	DebugLevel logrus.Level = logrus.DebugLevel
	ErrorLevel logrus.Level = logrus.ErrorLevel
	WarnLevel  logrus.Level = logrus.WarnLevel
	InfoLevel  logrus.Level = logrus.InfoLevel
	FatalLevel logrus.Level = logrus.FatalLevel
	PanicLevel logrus.Level = logrus.PanicLevel
)

// ContextLogger implements log.ContextLogger
type ContextLogger struct {
	*logrus.Logger
}

var _ log.ContextLogger = &ContextLogger{}

// WithContext adds context to log entry
func (cl *ContextLogger) WithContext(ctx log.M) log.ContextLogger {
	return &entry{cl.WithFields(logrus.Fields(ctx))}
}

type entry struct {
	*logrus.Entry
}

func (e *entry) WithContext(ctx log.M) log.ContextLogger {
	return &entry{e.WithFields(logrus.Fields(ctx))}
}

// New creates new ContextLogger
func New() *ContextLogger {
	return &ContextLogger{Logger: logrus.New()}
}
