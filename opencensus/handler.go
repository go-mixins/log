package opencensus

import (
	"context"
	"log/slog"

	"go.opencensus.io/trace"
)

type slogHandler struct {
	next slog.Handler
}

func (h *slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

func (h *slogHandler) Handle(ctx context.Context, record slog.Record) error {
	spanContext := trace.FromContext(ctx).SpanContext()
	return h.next.WithAttrs([]slog.Attr{
		slog.String("trace_id", spanContext.TraceID.String()),
		slog.String("span_id", spanContext.SpanID.String()),
	}).Handle(ctx, record)
}

func (h *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &slogHandler{next: h.next.WithAttrs(attrs)}
}

func (h *slogHandler) WithGroup(name string) slog.Handler {
	return &slogHandler{next: h.next.WithGroup(name)}
}

func Wrap(src slog.Handler) slog.Handler {
	return &slogHandler{next: src}
}
