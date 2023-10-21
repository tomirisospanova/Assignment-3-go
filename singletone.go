package main

import (
    "fmt"
    "os"
    "sync"
)

type Logger struct {
    fileName string
    file     *os.File
    mu       sync.Mutex
}

var instance *Logger
var once sync.Once

func GetLogger() *Logger {
    once.Do(func() {
        instance = &Logger{
            fileName: "app.log",
        }
        instance.openLogFile()
    })
    return instance
}

func (l *Logger) openLogFile() {
    file, err := os.OpenFile(l.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening log file: %v\n", err)
        return
    }
    l.file = file
}

func (l *Logger) Log(message string) {
    l.mu.Lock()
    defer l.mu.Unlock()

    if l.file != nil {
        _, err := fmt.Fprintln(l.file, message)
        if err != nil {
            fmt.Printf("Error writing to log file: %v\n", err)
        }
    }
}

func main() {
    logger := GetLogger()
    logger.Log("Запись в журнал: Событие 1")
    logger.Log("Запись в журнал: Событие 2")
}
