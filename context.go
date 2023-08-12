package log

import (
	"context"
	"log/slog"
)

type attrsKey struct{}

// WithAttrs injects logger attrs into context
func WithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	return context.WithValue(ctx, attrsKey{}, append(attrs, Attrs(ctx)...))
}

// Attrs returs logger attrs stored in context
func Attrs(ctx context.Context) []slog.Attr {
	attrs, _ := ctx.Value(attrsKey{}).([]slog.Attr)
	return attrs
}
