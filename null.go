package log

import (
	"log"
)

var defaultLogger = null{}

type null struct{}

func (null) Debug(...interface{}) {}

func (null) Debugf(string, ...interface{}) {}

func (null) Info(...interface{}) {}

func (null) Infof(string, ...interface{}) {}

func (null) Warn(...interface{}) {}

func (null) Warnf(string, ...interface{}) {}

func (null) Error(...interface{}) {}

func (null) Errorf(string, ...interface{}) {}

func (null) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (null) Fatalf(s string, v ...interface{}) {
	log.Fatalf(s, v...)
}

func (null) Panic(v ...interface{}) {
	log.Panic(v...)
}

func (null) Panicf(s string, v ...interface{}) {
	log.Panicf(s, v...)
}

func (null) WithContext(M) ContextLogger {
	return null{}
}
