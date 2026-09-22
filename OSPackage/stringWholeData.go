package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("textFile.txt")
	if err != nil {
		fmt.Println("Could not read file due to:", err)
		return
	}

	fmt.Println(string(data))
}
