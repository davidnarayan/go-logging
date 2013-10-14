// Package logging provides a simple leveled logging API

package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// ISO8601 timestamp
const tsFormat = "2006-01-02T15:04:05.999-0700"

// Log levels
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	STATS
)

// Log levels are used to control the level of verbosity in the logs
type Level int

// String representations of each log level (for display)
var levelStrings = map[Level]string{
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
	STATS: "STATS",
}

func (level Level) String() string {
	highestLevel := len(levelStrings) - 1

	if level < 0 || int(level) > highestLevel {
		return "UNKNOWN"
	}

	return levelStrings[level]
}

// A Logger writes out log messages
type Logger struct {
	Name   string
	Level  Level
	Writer io.Writer
}

// Set the minimum log level 
func (l *Logger) SetLevel(level Level) {
	l.Level = level
}

// Set the output destination
func (l *Logger) SetWriter(w io.Writer) {
	l.Writer = w
}

// Log a message
func (l *Logger) Log(level Level, format string, v ...interface{}) {
	if level < l.Level {
		return
	}

	buf := []byte(fmt.Sprintf("%s %s %s: %s\n", time.Now().Format(tsFormat),
		l.Name, level, fmt.Sprintf(format, v...)))

	l.Writer.Write(buf)
}

//-----------------------------------------------------------------------------
// Default logger

func newDefaultLogger() *Logger {
	// Use the filename as the default name
	var name string

	_, file, _, ok := runtime.Caller(2)

	if ok {
		name = filepath.Base(file)
	}

	l := &Logger{
		Name:   name,
		Level:  TRACE,
		Writer: os.Stderr,
	}

	return l
}

var std = newDefaultLogger()

func SetLevel(level Level) {
	std.SetLevel(level)
}

func SetWriter(w io.Writer) {
	std.SetWriter(w)
}

func Stats(format string, v ...interface{}) {
	std.Log(STATS, format, v...)
}

func Fatal(format string, v ...interface{}) {
	std.Log(FATAL, format, v...)
	os.Exit(1)
}

func Error(format string, v ...interface{}) {
	std.Log(ERROR, format, v...)
}

func Warn(format string, v ...interface{}) {
	std.Log(WARN, format, v...)
}

func Info(format string, v ...interface{}) {
	std.Log(INFO, format, v...)
}

func Debug(format string, v ...interface{}) {
	std.Log(DEBUG, format, v...)
}
