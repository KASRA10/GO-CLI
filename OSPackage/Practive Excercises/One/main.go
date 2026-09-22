package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("workspace/data/logs", 0755)
	if err != nil {
		fmt.Println("Error MKDIRALL:", err)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error Getwd:", err)
		return
	}

	fmt.Println(cwd)

	answer := "0755 For Managing Dir, While 0644 Mostly use for a file mode access"
	fmt.Println("\n" + answer + "\n")
}
