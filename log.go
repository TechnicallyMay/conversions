package main

import "fmt"

type Logger struct {
    PrintDebug bool
    PrintInfo bool
}

func (l *Logger) Debug(format string, args ...any) {
    if l.PrintDebug {
        l.print(format, args...)
    }
}

func (l *Logger) Info(format string, args ...any) {
    if l.PrintInfo {
        l.print(format, args...)
    }
}

func (l *Logger) print(format string, args ...any) {
    fmt.Printf(format, args...)
}
