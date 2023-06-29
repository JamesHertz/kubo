package util

import (
	"fmt"
	"io"
	logger "log"
	"os"
)

const loggerFileMode = os.O_APPEND|os.O_CREATE|os.O_WRONLY

func NewLogger(filename string) *logger.Logger{

	log_dir := os.Getenv("LOG_DIR")
	var out io.Writer = os.Stdout

	if log_dir != "" {
		full_file_name := log_dir + "/" + filename

		file, err := os.OpenFile(full_file_name, loggerFileMode, 0666)
		
		if err != nil {
			panic(fmt.Sprintf("Error opening file %s: %v", full_file_name, err))
		}
		out = file
	}

	return logger.New(out, "", logger.LstdFlags)
}