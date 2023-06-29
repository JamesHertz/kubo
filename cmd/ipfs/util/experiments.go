package util

import (
	"os"
	"fmt"
	logger "log"
)

const loggerFileMode = os.O_APPEND|os.O_CREATE|os.O_WRONLY

func NewLogger(filename string) *logger.Logger{

	log_dir := os.Getenv("LOG_DIR")

	if log_dir == "" {
		panic("Variable LOG_DIR not set...")
	}

	full_file_name := log_dir + "/" + filename

	file, err := os.OpenFile(full_file_name, loggerFileMode, 0666)
	
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %v", full_file_name, err))
	}

	return logger.New(file, "", logger.LstdFlags)
}