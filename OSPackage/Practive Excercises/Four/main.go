package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Start time:", start)

	fmt.Println("Please enter a directory path (e.g. ./FolderName or C:\\Windows\\Fonts):")

	var userPath string
	_, err := fmt.Scanln(&userPath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if userPath == "" {
		fmt.Println("Error: Path cannot be empty")
		return
	}

	absPath, err := filepath.Abs(userPath)
	if err != nil {
		fmt.Printf("Error converting path: %v\n", err)
		return
	}

	cleanPath := filepath.Clean(absPath)

	fmt.Printf("\nReading directory: %s\n", cleanPath)

	entries, err := os.ReadDir(cleanPath)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", cleanPath, err)
		return
	}

	fmt.Printf("=== Contents of %s ===\n", cleanPath)
	if len(entries) == 0 {
		fmt.Println("(empty directory)")
	} else {
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			isDirStr := "No"
			if info.IsDir() {
				isDirStr = "Yes"
			}

			fmt.Printf("Name: %s | Type: %s | Size: %d bytes | IsDir: %s | Last Modified: %v\n",
				entry.Name(),
				entry.Type(),
				info.Size(),
				isDirStr,
				info.ModTime().Format("2006-01-02 15:04:05"))
		}
	}

	fmt.Printf("\nElapsed time: %v\n", time.Since(start))
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
