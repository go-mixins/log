package log

import (
	"context"
	"log/slog"
)

type slogHandler struct {
	next slog.Handler
}

func (h *slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

func (h *slogHandler) Handle(ctx context.Context, record slog.Record) error {
	return h.next.WithAttrs(Attrs(ctx)).Handle(ctx, record)
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
