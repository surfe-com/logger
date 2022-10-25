package simple

import (
	"context"
	"log"

	"github.com/Leadjet/logger/logi"
)

type Logger struct{}

// WithContext does not work with Simple Logger
func (w *Logger) WithContext(ctx context.Context) logi.Logger {
	return w
}

func (w *Logger) Errorf(template string, err error, args ...interface{}) {
	args = append(args, err)
	log.Printf(template+". Error: %v", args...)
}

func (w *Logger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	log.Printf(msg+"%v %v ", keysAndValues, err)
}

func (w *Logger) Error(err error, args ...interface{}) {
	log.Printf("%v %v", err, args)
}

func (w *Logger) Warnf(template string, err error, args ...interface{}) {
	args = append(args, err)
	log.Printf(template+". Error: %v", args...)
}

func (w *Logger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	log.Printf(msg+"%v %v ", keysAndValues, err)
}

func (w *Logger) Warn(err error, args ...interface{}) {
	log.Printf("%v %v", err, args)
}

func (w *Logger) Infof(template string, args ...interface{}) {
	log.Printf(template, args...)
}

func (w *Logger) Infow(msg string, keysAndValues ...interface{}) {
	log.Printf(msg+"%v ", keysAndValues)
}

func (w *Logger) Info(args ...interface{}) {
	log.Printf("%v ", args...)
}

func (w *Logger) Debugf(template string, args ...interface{}) {
	log.Printf(template, args...)
}

func (w *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	log.Printf(msg+"%v ", keysAndValues)
}

func (w *Logger) Debug(args ...interface{}) {
	log.Printf("%v ", args)
}

func (w *Logger) Sync() {
	// Nothing to sync
}
