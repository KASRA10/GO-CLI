package logger

import (
	"error"
	"fmt"
	"log"
	"os"
	"time"

	"path.filepath"
)

func ELogger(message string) (*log.Logger, *os.File, error) {
	// step1: check directory for Log Folder
	entires, err := os.ReadDir("./")
	if err != nil {
		errors.New("Could Not Read Current Directory")
	}

	for _, entry := range entires {
	}
	// End Of Step1
}
