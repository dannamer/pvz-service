package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	With(ctx context.Context) Logger
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string, err error)
	Warn(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
}

type Zerologger struct {
	z zerolog.Logger
}

func New(pretty bool) *Zerologger {
	zerolog.TimeFieldFormat = time.RFC3339

	var z zerolog.Logger
	if pretty {
		z = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		}).With().Logger()
	} else {
		z = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Str("service", "pvz-service").
			Logger()
	}

	return &Zerologger{z: z}
}

func (l *Zerologger) With(ctx context.Context) Logger {
	fields := l.z.With()

	if rid, ok := ctx.Value("request_id").(string); ok {
		fields = fields.Str("request_id", rid)
	}
	if method, ok := ctx.Value("method").(string); ok {
		fields = fields.Str("method", method)
	}
	if path, ok := ctx.Value("path").(string); ok {
		fields = fields.Str("path", path)
	}
	if uid, ok := ctx.Value("user_id").(string); ok {
		fields = fields.Str("user_id", uid)
	}

	return &Zerologger{z: fields.Logger()}
}

func (l *Zerologger) Info(ctx context.Context, msg string) {
	l.z.Info().Msg(msg)
}

func (l *Zerologger) Error(ctx context.Context, msg string, err error) {
	l.z.Error().Err(err).Msg(msg)
}

func (l *Zerologger) Warn(ctx context.Context, msg string) {
	l.z.Warn().Msg(msg)
}

func (l *Zerologger) Debug(ctx context.Context, msg string) {
	l.z.Debug().Msg(msg)
}
