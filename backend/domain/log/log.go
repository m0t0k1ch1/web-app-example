package log

import (
	"log/slog"
	"os"
	"strconv"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(strconv.FormatInt(a.Value.Time().Unix(), 10))
			}

			return a
		},
	}))
}
