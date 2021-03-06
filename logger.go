package main

import "fmt"
import "os"

type Logger struct {
	name string
	level int
}

const (
	NONE = iota
	FATAL
	ERROR
	INFO
	DEBUG
)

func NewLogger(name string, level int) *Logger {
	/*if len(name) > 6 {
		name = name[0:6]
	}*/
	return &Logger{name: name, level: level}
}

func levelString(level int) string {
	switch level {
		case FATAL:
			return "FATAL"
		case ERROR:
			return "ERROR"
		case DEBUG:
			return "DEBUG"
		case INFO:
			return " INFO"
	}
	return ""
}

func (log *Logger) print(dev *os.File, level int, msg string) {
	if log.level >= level {
		//fmt.Fprintf(dev, "%s: %6s: %s", levelString(level), log.name, msg)
		fmt.Fprintf(dev, "%s%s", log.name, msg)
		//fmt.Fprintf(dev, "%s", msg)
	}
}

func (log *Logger) Fatal(msg ...interface{}) {
	buf := fmt.Sprintln(msg...)
	log.print(os.Stderr, FATAL, buf)
	os.Exit(1)
}

func (log *Logger) Error(msg ...interface{}) {
	buf := fmt.Sprintln(msg...)
	log.print(os.Stderr, ERROR, buf)
}

func (log *Logger) Info(msg ...interface{}) {
	buf := fmt.Sprintln(msg...)
	log.print(os.Stderr, INFO, buf)
}

func (log *Logger) Debug(msg ...interface{}) {
	buf := fmt.Sprintln(msg...)
	log.print(os.Stderr, DEBUG, buf)
}

func (log *Logger) Debugf(format string, msg ...interface{}) {
	buf := fmt.Sprintf(format, msg...)
	log.print(os.Stderr, DEBUG, buf)
}

