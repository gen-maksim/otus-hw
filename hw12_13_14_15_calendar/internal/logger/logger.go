package logger

import (
	"log"
)

type Logger struct {
	Level int
}

const (
	lvlInfo = iota
	lvlWarn
	lvlDebug
	lvlError
)

var lvlMap = map[string]int{
	"info":  lvlInfo,
	"warn":  lvlWarn,
	"debug": lvlDebug,
	"error": lvlError,
}

func New(level string) *Logger {
	l, ok := lvlMap[level]
	if !ok {
		log.Fatal("no logger for level ", level)
	}

	return &Logger{Level: l}
}

func (l Logger) Info(msg string) {
	if !l.checkLvl(lvlInfo) {
		return
	}

	log.Printf("Info: %v\n", msg)
}

func (l Logger) Warn(msg string) {
	if !l.checkLvl(lvlWarn) {
		return
	}

	log.Printf("Warn: %v\n", msg)
}

func (l Logger) Debug(msg string) {
	if !l.checkLvl(lvlDebug) {
		return
	}

	log.Printf("Debug: %v\n", msg)
}

func (l Logger) Error(msg string) {
	if !l.checkLvl(lvlError) {
		return
	}

	log.Printf("Error: %v\n", msg)
}

func (l Logger) checkLvl(lvl int) bool {
	return lvl >= l.Level
}
