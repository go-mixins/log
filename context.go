package log

import "context"

type loggerKeyType struct{}

var loggerKey loggerKeyType

// With injects logger into context
func With(ctx context.Context, logger ContextLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// Get logger from the context. Returned logger is always safe to use.
func Get(ctx context.Context) ContextLogger {
	if res, ok := ctx.Value(loggerKey).(ContextLogger); ok {
		return res
	}
	return defaultLogger
}
