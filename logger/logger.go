package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	LevelDebug int = iota
	LevelVerbose
	LevelInfo
	LevelWarn
	LevelError
)

type Config struct {
	Level     int
	UseErrout bool
	Stdout    io.Writer
	Errout    io.Writer
}

func (config *Config) SetLevel(level int) *Config {
	config.Level = level
	return config
}
func (config *Config) WithErrout(v bool) *Config {
	config.UseErrout = v
	return config
}
func (config *Config) SetErrout(out io.Writer) *Config {
	config.Errout = out
	return config
}
func (config *Config) SetStdout(out io.Writer) *Config {
	config.Stdout = out
	return config
}

type Logger struct {
	Config Config
}

func (logger *Logger) Parse(str string, level int, args []any) string {
	str = fmt.Sprintf(str, args...)

	now := time.Now().Format("2006-01-02 15:04:05.000")
	levelString := LevelString(level)
	return fmt.Sprintf("[%s] [%s] %s", now, levelString, str)
}

func (logger *Logger) Print(str string, level int, args []any) {
	out := logger.Config.Stdout
	if level < logger.Config.Level {
		return
	}
	if level >= LevelError && logger.Config.UseErrout {
		out = logger.Config.Errout
	}

	parsed := logger.Parse(str, level, args)

	fmt.Fprintln(out, parsed)
}

func LevelString(level int) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelVerbose:
		return "VERBOSE"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return ""
	}
}

func (logger *Logger) LevelString() string {
	return LevelString(logger.Config.Level)
}

func NewLogger() *Logger {
	config := &Config{
		Level:     LevelInfo,
		UseErrout: false,
		Stdout:    os.Stdout,
		Errout:    os.Stderr,
	}

	return &Logger{
		Config: *config,
	}
}

func (logger *Logger) Debug(str string, args ...any) {
	logger.Print(str, LevelDebug, args)
}
func (logger *Logger) Verbose(str string, args ...any) {
	logger.Print(str, LevelVerbose, args)
}
func (logger *Logger) Info(str string, args ...any) {
	logger.Print(str, LevelInfo, args)
}
func (logger *Logger) Warn(str string, args ...any) {
	logger.Print(str, LevelWarn, args)
}
func (logger *Logger) Error(str string, args ...any) {
	logger.Print(str, LevelError, args)
}
