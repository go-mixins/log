// Package graylog adds to github.com/go-mixins/logrus.ContextLogger the
// ability to log entries to Graylog
//
package graylog

import (
	"github.com/go-mixins/log/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

// NewAsyncGraylogHook copied here to reduce import overhead
var NewAsyncGraylogHook = graylog.NewAsyncGraylogHook

// Some useful constants from Logrus to reduce import clutter
const (
	DebugLevel = logrus.DebugLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

type ContextLogger struct {
	*logrus.ContextLogger
	hooks []*graylog.GraylogHook
}

func (c *ContextLogger) Flush() {
	for _, h := range c.hooks {
		h.Flush()
	}
}

// New creates new ContextLogger with Logrus side-logging to specified URIs
func New(graylogURIs ...string) *ContextLogger {
	res := &ContextLogger{
		ContextLogger: logrus.New(),
	}
	for _, uri := range graylogURIs {
		if uri == "" {
			continue
		}
		hook := graylog.NewAsyncGraylogHook(uri, nil)
		res.Hooks.Add(hook)
		res.hooks = append(res.hooks, hook)
	}
	return res
}
