package logger

import (
	"log"
	"os"
)

type Logger interface {
	Printf(format string, v ...any)
}

func New() *log.Logger {
	return log.New(os.Stdout, "task59", log.Flags())
}
