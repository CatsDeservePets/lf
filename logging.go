package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type logLevel int

const (
	logDebug logLevel = iota
	logInfo
	logWarn
	logError
)

var prefixes = [...]string{
	"[DEBUG] ",
	"[INFO]  ",
	"[WARN]  ",
	"[ERROR] ",
}

func (l logLevel) String() string {
	switch l {
	case logDebug:
		return "debug"
	case logInfo:
		return "info"
	case logWarn:
		return "warn"
	case logError:
		return "error"
	default:
		return fmt.Sprintf("logLevel(%d)", l)
	}
}

func (l *logLevel) Set(s string) error {
	switch strings.ToLower(s) {
	case "debug":
		*l = logDebug
	case "info":
		*l = logInfo
	case "warn":
		*l = logWarn
	case "error":
		*l = logError
	default:
		return errors.New(`valid levels are "debug", "info", "warn", or "error"`)
	}
	return nil
}

func logf(level logLevel, format string, args ...any) {
	if level < gLogLevel {
		return
	}

	log.Printf(prefixes[level]+format, args...)
}

func debugf(format string, args ...any) {
	logf(logDebug, format, args...)
}

func infof(format string, args ...any) {
	logf(logInfo, format, args...)
}

func warnf(format string, args ...any) {
	logf(logWarn, format, args...)
}

func errorf(format string, args ...any) {
	logf(logError, format, args...)
}
