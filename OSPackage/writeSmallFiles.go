package main

import (
	"fmt"
	"os"
)

func main() {
	message := []byte("\nThis Is My New MESSAGE\n")

	err := os.WriteFile("textFile.txt", message, 0644)
	if err != nil {
		fmt.Println("Write Error:", err)
		return
	}

	data, err := os.ReadFile("textFile.txt")
	if err != nil {
		fmt.Println("Read Error:", err)
		return
	}

	fmt.Println(string(data))

	/*
			0644 is a Unix-style permission mode: owner can read/write; group and others can read. On Windows, permission behavior is more limited, so do not depend on Unix permission details for cross-platform security.

			Important: os.WriteFile replaces old content. It is not an append operation, and a failure during its multi-step write can leave partial content. It delete old COntent and not just append, replace with new content

			if do not use 0644 for windows:
			error:
			.\writeSmallFiles.go:11:38: not enough arguments in call to os.WriteFile
	        have (string, []byte)
	        want (string, []byte, os.FileMode)
	*/
}
