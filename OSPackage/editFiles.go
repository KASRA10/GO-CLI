package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("textFileTwo.txt")
	if err != nil {
		fmt.Println("File Creating Error:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile("textFileTwo.txt")
	if err != nil {
		fmt.Println("File Reading Error:", err)
		return
	}
	fmt.Println(string(data))

	newMessage := []byte("this is first message\n")

	err = os.WriteFile("textFileTwo.txt", newMessage, 0644)
	if err != nil {
		fmt.Println("File Writing Error:", err)
		return
	}

	data, err = os.ReadFile("textFileTwo.txt")
	if err != nil {
		fmt.Println("File Reading Error:", err)
		return
	}
	fmt.Println(string(data))

	file, err = os.Create("textFileTwo.txt")
	if err != nil {
		fmt.Println("File Creating Error:", err)
		return
	}
	defer file.Close()

	data, err = os.ReadFile("textFileTwo.txt")
	if err != nil {
		fmt.Println("File Reading Error:", err)
		return
	}
	fmt.Println(string(data))

	/*
		Be careful: os.Create("output.txt") truncates an existing file to zero bytes. Use it only when replacing the file is intended.
	*/
}
