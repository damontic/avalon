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
	Source  string `json:"source"`
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

func (l Logger) Debug(source string, message string) {
	l.print(DEBUG, source, message)
}

func (l Logger) Debugf(source string, format string, params ...interface{}) {
	message := fmt.Sprintf(format, params...)
	l.print(DEBUG, source, message)
}

func (l Logger) Info(source string, message string) {
	l.print(INFO, source, message)
}

func (l Logger) Infof(source string, format string, params ...interface{}) {
	message := fmt.Sprintf(format, params...)
	l.print(INFO, source, message)
}

func (l Logger) Warn(source string, message string) {
	l.print(WARN, source, message)
}

func (l Logger) Warnf(source string, format string, params ...interface{}) {
	message := fmt.Sprintf(format, params...)
	l.print(WARN, source, message)
}

func (l Logger) Error(source string, message string) {
	l.print(ERROR, source, message)
}

func (l Logger) Errorf(source string, format string, params ...interface{}) {
	message := fmt.Sprintf(format, params...)
	l.print(ERROR, source, message)
}

func (l Logger) Fatal(source string, message string) {
	l.print(FATAL, source, message)
	panic(message)
}

func (l Logger) Fatalf(source string, format string, params ...interface{}) {
	message := fmt.Sprintf(format, params...)
	l.print(FATAL, source, message)
	panic(message)
}

func (l Logger) print(level VerbosityLevel, source string, message string) error {
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
		Source:  source,
		Message: message,
	}
	b, err := json.Marshal(logEvent)
	if err != nil {
		return fmt.Errorf("Error Marshalling logEvent")
	}
	l.logger.Printf(string(b[:]))
	fmt.Print(l.buffer)
	l.buffer.Reset()
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
