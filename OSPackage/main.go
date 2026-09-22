package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("textFile2.txt")
	if err != nil {
		fmt.Println(err)
	} // open textFile2.txt: The system cannot find the file specified.

	file, err2 := os.Open("textFile.txt")
	if err2 != nil {
		fmt.Println(err2)
	}

	data := make([]byte, 100)
	count, err3 := file.Read(data)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("Data Is:", data)
	fmt.Println("Length Is:", len(data))
	fmt.Println("Count Is:", count)
	fmt.Printf("Read %d bytes: %q\n", count, data[:count])
	/*
		result:
		open textFile2.txt: The system cannot find the file specified.
		Data Is: [76 111 114 101 109 32 105 112 115 117 109 32 100 111 108 111 114 32 115 105 116 32 97 109 101 116 32 99 111 110 115 10199 116 101 116 117 114 32 97 100 105 112 105 115 99 105 110 103 32 101 108 105 116 46 32 81 117 105 115 113 117 101 32 102 97 11799 105 98 117 115 32 101 120 32 115 97 112 105 101 110 32 118 105 116 97 101 32 112 101 108 108 101 110 116 101 115 113 117]
		Length Is: 100
		Count Is: 100
		Read 100 bytes: "Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesqu"
	*/

	defer file.Close()
	// It closes the operating-system file resource when main finishes.
}
