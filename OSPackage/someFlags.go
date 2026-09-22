package main

import (
	"fmt"
	"os"
)

/*
The first flag must choose access mode; additional flags are combined with |.

Flag	Meaning
os.O_RDONLY	Read only
os.O_WRONLY	Write only
os.O_RDWR	Read and write
os.O_CREATE	Create file if missing
os.O_APPEND	Always write at the end
os.O_TRUNC	Empty the file when opening
os.O_EXCL	With O_CREATE, fail if file already exists

For logs, the common combination is:

go
os.O_WRONLY | os.O_CREATE | os.O_APPEND
For a new file that must not overwrite an existing one, use:

go
os.O_WRONLY | os.O_CREATE | os.O_EXCL
*/

func main() {
	file, err := os.OpenFile("textFileThree.txt", os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	message := "lorem lorem lorem lorem lorem lorem lorem lorem"
	err = os.WriteFile("textFileThree.txt", []byte(message), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := os.ReadFile("textFileThree.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
