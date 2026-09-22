package main

import (
	"fmt"
	"os"
)

func main() {
	folderArray := [3]string{
		"src/uploades/images",
		"src/uploades/icons",
		"src/uploades/fonts",
	}

	for _, val := range folderArray {

		err := os.MkdirAll(val, 0755)
		if err != nil {
			fmt.Println("Directory creation error:", err)
			return
		}
	}

	entries, err := os.ReadDir("./src/uploades")
	if err != nil {
		fmt.Println("ReadDir error:", err)
		return
	}

	for index, entry := range entries {

		fmt.Println("index:", index)
		fmt.Println("Name:", entry.Name())
		fmt.Println("Type:", entry.Type())
		fmt.Println("IsDir:", entry.IsDir())
	}

	fmt.Println("-------Second Part-------")

	entries, err = os.ReadDir("./")
	if err != nil {
		fmt.Println("ReadDir error:", err)
		return
	}

	file, err := os.OpenFile(
		"information.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println("Open Error:", err)
		return
	}

	defer file.Close()

	var i int = 1

	for _, entry := range entries {

		info, err := entry.Info()
		if err != nil {
			fmt.Println("Info Error:", err)
			continue
		}
		fmt.Println("-------Entry-------", i)
		fmt.Println("Name:", entry.Name())
		fmt.Println("Bytes:", info.Size())
		fmt.Println("Modified:", info.ModTime())
		fmt.Println("IsDir:", info.IsDir())
		fmt.Println("Mode:", info.Mode())
		fmt.Println("Sys:", info.Sys())
		fmt.Println("------- End Of Entry-------", i)

		message := fmt.Sprintf(
			"------- Entry ------- %d\n"+
				"Name: %s\n"+
				"Bytes: %d\n"+
				"Modified: %s\n"+
				"IsDir: %v\n"+
				"Mode: %s\n"+
				"Sys: %v\n"+
				"------- End Of Entry ------- %d\n\n",
			i,
			entry.Name(),
			info.Size(),
			info.ModTime(),
			info.IsDir(),
			info.Mode(),
			info.Sys(),
			i,
		)
		_, err = file.WriteString(message)
		if err != nil {
			fmt.Println("Write Error:", err)
			return
		}

		i++
	}

	fmt.Println("********************************")
	fmt.Println("-------Third Part-------")
	fmt.Println("********************************")

	data, err := os.ReadFile("information.txt")
	if err != nil {
		fmt.Println("Reading Error:", err)
		return
	}

	fmt.Println(string(data))
}

/*
Placeholder,Example,What it does
%s,%s,String
%d / %v,%d or %v,Integer / any number
%f,%f,Float number
%v,%v,Any value (best all-rounder)
%+v,%+v,Like %v but shows struct fields
%T,%T,Type of the value
*/
