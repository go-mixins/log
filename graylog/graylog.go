// Package graylog adds to github.com/go-mixins/logrus.ContextLogger the
// ability to log entries to Graylog
//
package graylog

import (
	"github.com/go-mixins/log/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

// Some useful constants from Logrus to reduce import clutter
const (
	DebugLevel = logrus.DebugLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

// New creates new ContextLogger with Logrus side-logging to specified URLs
func New(graylogURI string) *logrus.ContextLogger {
	res := logrus.New()
	res.Hooks.Add(graylog.NewAsyncGraylogHook(graylogURI, nil))
	return res
}
