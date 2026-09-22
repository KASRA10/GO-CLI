package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("src/uploads/images", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll("src/uploads/icons", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll("src/uploads/fonts", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	entries, err := os.ReadDir("./src/uploads")
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, entry := range entries {
		fmt.Println("index is:", index, "<=>", "directory:", entry.Name(), "folder?", entry.IsDir(), "The Type Is:", entry.Type())
	}

	fmt.Println("-----------------------------------------")

	entries, err = os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		info, err := os.Stat(entry.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		if entry.IsDir() != true {
			fmt.Println("Bytes:", info.Size())
			fmt.Println("Last Modified:", info.ModTime())
			fmt.Println("Is A Folder:", info.IsDir())
			fmt.Println("Mode", info.Mode())
			fmt.Println("PlatForm Specification:", info.Sys())
			fmt.Println("-----------------------------------------")
		}
	}
}
