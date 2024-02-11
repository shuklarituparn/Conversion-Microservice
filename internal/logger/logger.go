package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var Log *os.File

func InitLog() *os.File {
	logFileName := fmt.Sprintf("%s_%s.log", time.Now().Format("2006-01-02"), time.Now())
	logFilePath := filepath.Join("../../internal/logs", logFileName)
	f, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return nil
	}

	Log = f

	return Log
}
