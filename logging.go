// Package logging provides a simple leveled logging API

package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
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
	mu     sync.Mutex
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

// Set the logger name
func (l *Logger) SetName(name string) {
	l.Name = name
}

// Log a message
func (l *Logger) Log(level Level, format string, v ...interface{}) {
	if level < l.Level {
		return
	}

	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()
	name := l.Name
	var trace string

	// Trace info
	if l.Level == TRACE {
		_, file, line, ok := runtime.Caller(2)

		if !ok {
			file = "???"
			line = 0
		} else {
			slash := strings.LastIndex(file, "/")
			if slash >= 0 {
				file = file[slash+1:]
			}
		}

		var trc []string
		trc = append(trc, "")

		if !strings.HasPrefix(name, file) {
			trc = append(trc, file)
		}

		trc = append(trc, strconv.Itoa(line))
		trace = strings.Join(trc, ":")
	}

	buf := []byte(fmt.Sprintf("%s %s%s %s: %s\n", now.Format(tsFormat), name,
		trace, level, fmt.Sprintf(format, v...)))

	l.Writer.Write(buf)
}

//-----------------------------------------------------------------------------
// Default logger

func newDefaultLogger() *Logger {
	// Use the filename and pid as the default name
	var name string

	_, file, _, ok := runtime.Caller(2)

	if ok {
		name = fmt.Sprintf("%s[%d]", filepath.Base(file), os.Getpid())
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

func Trace(format string, v ...interface{}) {
	std.Log(TRACE, format, v...)
}
