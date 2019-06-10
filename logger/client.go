package logger

import (
	"fmt"
	"os"
	"time"
)

func Info(msg ...interface{}) {
	fmt.Fprintln(os.Stdout,time.Now().String()[:19], msg)
}

func Error(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, time.Now().String()[:19], msg)
}
