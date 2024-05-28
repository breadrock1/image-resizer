package local

import (
	"fmt"
	"log"
	"os"

	"image-resize-service/internal/config"
)

type Logger struct {
	InfoLogger *log.Logger
	WarnLogger *log.Logger
	ErrLogger  *log.Logger
	loggerFile *os.File
}

func New(cfg *config.LoggerConfig) Logger {
	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	if !cfg.EnableFileLog {
		return buildDefaultLogger(logFlags)
	}

	osFlags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(cfg.FilePath, osFlags, 0o666)
	if err != nil {
		fmt.Printf("faile to init log: %s", err)
		return buildDefaultLogger(logFlags)
	}

	return Logger{
		InfoLogger: log.New(file, "INFO: ", logFlags),
		WarnLogger: log.New(file, "WARN: ", logFlags),
		ErrLogger:  log.New(file, "ERROR: ", logFlags),
		loggerFile: file,
	}
}

func buildDefaultLogger(logFlags int) Logger {
	return Logger{
		InfoLogger: log.New(os.Stdout, "INFO: ", logFlags),
		WarnLogger: log.New(os.Stdout, "WARN: ", logFlags),
		ErrLogger:  log.New(os.Stdout, "ERROR: ", logFlags),
	}
}

func (l *Logger) Info(msg ...string) {
	l.InfoLogger.Println(msg)
}

func (l *Logger) Warn(msg ...string) {
	l.WarnLogger.Println(msg)
}

func (l *Logger) Error(msg ...string) {
	l.ErrLogger.Println(msg)
}
