package logger

import (
	"log"

	"github.com/avila-r/moon/config"
)

type DebugLogger struct {
	Active bool
}

func Debugger() *DebugLogger {
	return &DebugLogger{
		Active: config.Get().Advanced.Debug,
	}
}

func (d *DebugLogger) Log(v ...any) {
	if !d.Active {
		return
	}

	log.Print(v...)
}

func (d *DebugLogger) Logf(fmt string, v ...any) {
	if !d.Active {
		return
	}

	log.Printf(fmt, v...)
}
