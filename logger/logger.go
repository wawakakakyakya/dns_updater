package logger

import (
	"dns_updater/config"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger *zerolog.Logger
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) InfoF(msg string, a ...any) {
	l.Info(fmt.Sprintf(msg, a...))
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *Logger) WarnF(msg string, a ...any) {
	l.Warn(fmt.Sprintf(msg, a...))
}

func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *Logger) ErrorF(msg string, a ...any) {
	l.Error(fmt.Sprintf(msg, a...))
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *Logger) DebugF(msg string, a ...any) {
	l.Debug(fmt.Sprintf(msg, a...))
}

func (l *Logger) Child(process string) *Logger {
	logger := l.logger.With().Str("process", process).Logger()
	return &Logger{logger: &logger}
}

func NewLogger(process string, cfg *config.LogConfig) *Logger {
	var logger zerolog.Logger
	rotateWriter := &lumberjack.Logger{
		Filename:   cfg.Path,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,   //days
		Compress:   cfg.Compress, // disabled by default
	}
	stdoutWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
	writer := io.MultiWriter(stdoutWriter, rotateWriter)
	logger = zerolog.New(writer).Level(zerolog.Level(cfg.Level)).With().
		Timestamp().
		Str("process", process).
		Logger()

	return &Logger{logger: &logger}
}
