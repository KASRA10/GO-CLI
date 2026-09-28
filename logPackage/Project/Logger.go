package logger

import (
	"errors"
	"fmt"
	"os"
)

func ELogger(message string) {
	// Global variables
	currentPath, err := os.Getwd()
	if err != nil {
		err = errors.New("Cannot Access To Current Path")
		fmt.Println(err)
	}
	fmt.Printf("CWD: %v\n", currentPath)

	// Step1: Check Log Directory
	entires, err := os.ReadDir(currentPath)
	if err != nil {
		err = errors.New("Cannot Read Directory")
		fmt.Println(err)
	}

	fmt.Printf("%T\n", entires)
	// End Of Step1
}
