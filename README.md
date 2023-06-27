# Package log v2

This is a total rewrite of the original package. Its only reason for existence
is some features that are not available in current (at the time of writing)
implementation of slog package.

## Using the package

Create JSON, Text, or any other handler as usual, then wrap it using `log.Wrap`:

```go
logger := log.Wrap(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
}))
```

After that you can use logger directly or replace the global logger with wrapped
version like this:

```go
slog.SetDefault(slog.New(logger))
```

Next, when there's some context in the application, like incoming HTTP request
you can add log attributes to that context. E.g. middleware to add request ID to
HTTP request:

```go
func AddIDToContext(src http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := log.WithAttrs(r.Context(), []slog.Attr{
			slog.String("request_id", uuid.NewString()),
		})
		src.ServeHTTP(w, r.WithContext(ctx))
	}
}
```

Finally, use context-aware logger methods, like `slog.InfoCtx` to output log
entries with common attributes added.
