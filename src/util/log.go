package util

import (
	"fmt"
	"log"
	"os"
)

func base(format string, level string, v ...interface{}) {
	log.Printf("[%s]\t- %s \n", level, fmt.Sprintf(format, v...))
}

func DEBUG(format string, v ...interface{}) {
	base(format, "DEBUG", v...)
}

func ERROR(format string, v ...interface{}) {
	base(format, "ERROR", v...)
}

func INFO(format string, v ...interface{}) {
	base(format, "INFO", v...)
}

func FATAL(format string, v ...interface{}) {
	base(format, "FATAL", v...)
	os.Exit(1)
}
