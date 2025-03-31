package logger

import (
	"context"
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	zlog *zerolog.Logger
}

func New() *Logger {
	cw := zerolog.NewConsoleWriter()
	mw := io.MultiWriter(cw)
	zlog := zerolog.New(mw).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
	return &Logger{
		zlog: &zlog,
	}
}

func (l *Logger) Trace(ctx context.Context, message string, fields ...any) {
	l.zlog.Trace().Fields(fields).Msg(message)
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...any) {
	l.zlog.Debug().Fields(fields).Msg(message)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...any) {
	l.zlog.Info().Fields(fields).Msg(message)
}

func (l *Logger) Warn(ctx context.Context, message string, fields ...any) {
	l.zlog.Warn().Fields(fields).Msg(message)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...any) {
	l.zlog.Error().Fields(fields).Msg(message)
}

func (l *Logger) Fatal(ctx context.Context, message string, fields ...any) {
	l.zlog.Fatal().Fields(fields).Msg(message)
}

func (l *Logger) Panic(ctx context.Context, message string, fields ...any) {
	l.zlog.Panic().Fields(fields).Msg(message)
}
