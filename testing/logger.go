package testing

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-mixins/log"
)

type Logger struct {
	t     *testing.T
	entry log.M
	level int
}

var _ log.ContextLogger = (*Logger)(nil)

func (l *Logger) format(lvl int, f string, v ...interface{}) string {
	data := log.M{
		"message": fmt.Sprintf(f, v...),
	}
	switch lvl {
	case 0:
		data["level"] = "debug"
	case 1:
		data["level"] = "info"
	case 2:
		data["level"] = "warn"
	case 3:
		data["level"] = "fatal"
	case 4:
		data["level"] = "panic"
	}
	jd, _ := json.Marshal(merge(l.entry, data))
	return string(jd)
}

func New(t *testing.T) *Logger {
	return &Logger{
		t: t,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level > 0 {
		return
	}
	l.t.Log(l.format(0, "%s", fmt.Sprint(v...)))
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	if l.level > 0 {
		return
	}
	l.t.Log(l.format(0, f, v...))
}

func (l *Logger) Info(v ...interface{}) {
	if l.level > 1 {
		return
	}
	l.t.Log(l.format(1, "%s", fmt.Sprint(v...)))
}

func (l *Logger) Infof(f string, v ...interface{}) {
	if l.level > 1 {
		return
	}
	l.t.Log(l.format(1, f, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level > 2 {
		return
	}
	l.t.Log(l.format(2, "%s", fmt.Sprint(v...)))
}

func (l *Logger) Warnf(f string, v ...interface{}) {
	if l.level > 2 {
		return
	}
	l.t.Log(l.format(2, f, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.t.Log(l.format(3, "%s", fmt.Sprint(v...)))
}

func (l *Logger) Errorf(f string, v ...interface{}) {
	l.t.Log(l.format(3, f, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.t.Fatal(l.format(4, "%s", fmt.Sprint(v...)))
}

func (l *Logger) Fatalf(f string, v ...interface{}) {
	l.t.Fatal(l.format(4, f, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	panic(l.format(5, "%s", fmt.Sprint(v...)))
}

func (l *Logger) Panicf(f string, v ...interface{}) {
	panic(l.format(5, f, v...))
}

func merge(a, b log.M) log.M {
	data := make(log.M)
	for k, v := range a {
		data[k] = v
	}
	for k, v := range b {
		for {
			if _, ok := data[k]; !ok {
				break
			}
			k = "_" + k
		}
		data[k] = v
	}
	return data
}

func (l *Logger) WithContext(entry log.M) log.ContextLogger {
	return &Logger{
		t:     l.t,
		entry: merge(l.entry, entry),
		level: l.level,
	}
}
