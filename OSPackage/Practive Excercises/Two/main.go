package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("workspace/data", 0755)
	if err != nil {
		fmt.Println("Error Creating Directory:", err)
		return
	}

	var title string

	fmt.Println("Please Enter Your Journal Title Here:")
	fmt.Scanln(&title)

	filePath := "workspace/data/" + title + ".txt"

	file, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		fmt.Println("Error Opening File:", err)
		return
	}

	defer file.Close()

	var journalEntryLine string

	fmt.Println("Please Enter Your Journal Entry Here:")
	fmt.Scanln(&journalEntryLine)

	_, err = file.WriteString(journalEntryLine + "\n")
	if err != nil {
		fmt.Println("Error Writing File String:", err)
		return
	}

	fmt.Println("Journal Entry Added Successfully!")
	fmt.Println("Press Any Key And Enter To Quit...")

	var quit string
	fmt.Scan(&quit)
}
