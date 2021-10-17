package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type VerbosityLevel int

const (
	OFF VerbosityLevel = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)

type Logger struct {
	Level  VerbosityLevel
	logger *log.Logger
	buffer *bytes.Buffer
	useUtc bool
}

type logEvent struct {
	Date    string `json:"date"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

func New(level VerbosityLevel, useUtc bool) (*Logger, error) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Lmsgprefix)
	return &Logger{
		Level:  level,
		logger: logger,
		buffer: &buf,
		useUtc: useUtc,
	}, nil
}

func (l Logger) Info(message string) {
	l.print(INFO, message)
}

func (l Logger) Warn(message string) {
	l.print(WARN, message)
}

func (l Logger) Error(message string) {
	l.print(ERROR, message)
}

func (l Logger) Fatal(message string) {
	l.print(FATAL, message)
}

func (l Logger) print(level VerbosityLevel, message string) error {
	if l.Level < level {
		return nil
	}
	var t string
	if l.useUtc {
		t = time.Now().UTC().Format(time.RFC3339)
	} else {
		t = time.Now().Format(time.RFC3339)
	}

	levelName := getLoggerName(level)
	logEvent := logEvent{
		Level:   levelName,
		Date:    t,
		Message: message,
	}
	b, err := json.Marshal(logEvent)
	if err != nil {
		return fmt.Errorf("Error Marshalling logEvent")
	}
	l.logger.Printf(string(b[:]))
	fmt.Print(l.buffer)
	return nil
}

func getLoggerName(level VerbosityLevel) string {
	switch level {
	case ALL:
		return "ALL"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case OFF:
		return "OFF"
	}

	if level < 0 {
		return "OFF"
	} else {
		return "ALL"
	}
}
