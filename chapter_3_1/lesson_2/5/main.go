package main // Тема 1/6: ООП → Урок 2/3 "Эмбеддинг"
// https://practicum.yandex.ru/trainer/go-basics/lesson/0130e521-80e9-4a8b-8faa-99db45d4a89a/

import (
	"log"
	"os"
)

// В стандартной библиотеке Go уже есть логгер — в пакете log. У этой реализации отсутствует параметр уровня лога (log level), то есть для вывода есть только метод log.Println и log.Printf. Предлагаем расширить этот объект следующими методами:
// SetLogLevel(logLvl LogLevel);
// Infoln(msg string);
// Warnln(msg string);
// Errorln(msg string).

type LogLevel int

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true

	default:
		return false
	}
}

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger // встраиваем стандартный логгер

	logLevel LogLevel
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)

	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
}
