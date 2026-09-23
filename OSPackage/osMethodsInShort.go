package main

import (
	"fmt"
	"os"
)

func main() {
	sliceOfPaths := []string{
		"structure/folder1/A",
		"structure/folder1/B",
		"structure/folder1/C",
	}

	for _, value := range sliceOfPaths {
		err := os.MkdirAll(value, 0755)
		if err != nil {
			fmt.Printf("%q\n", err)
		}
	}

	entries, err := os.ReadDir("./structure/folder1")
	if err != nil {
		fmt.Printf("%q\n", err)
	}

	i := 1
	for _, entry := range entries {
		fmt.Println(i)
		fmt.Println(entry.Name())
		fmt.Println(entry.Type())
		fmt.Println(entry.IsDir())
		fmt.Println(i, "\n")
		i++
	}

	file, err := os.OpenFile("newNote.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("%q\n", err)
	}
	defer file.Close()

	err = file.Close()
	if err != nil {
		fmt.Printf("%q\n", err)
	}

	err = os.Rename("./newNote.txt", "./structure/folder1/A/newNote.txt")
	if err != nil {
		fmt.Printf("%q\n", err)
	}

	entries, err = os.ReadDir("./")
	if err != nil {
		fmt.Printf("%q\n", err)
	}

	file, err = os.OpenFile("./structure/folder1/A/newNote.txt", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("%q\n", err)
	}

	j := 1
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("%q\n", err)
		}

		fmt.Println("-------Entry-------", j)
		fmt.Println("Name:", entry.Name())
		fmt.Println("Bytes:", info.Size())
		fmt.Println("Modified:", info.ModTime())
		fmt.Println("IsDir:", info.IsDir())
		fmt.Println("Mode:", info.Mode())
		fmt.Println("Sys:", info.Sys())
		fmt.Println("------- End Of Entry-------\n", j)

		textMessage := fmt.Sprintf(
			"-------Entry-------%d\n"+
				"Name: %s\n"+
				"Bytes: %d\n"+
				"Modified: %s\n"+
				"IsDir: %v\n"+
				"Mode: %s\n"+
				"Sys: %v\n"+
				"------- End Of Entry ------- %d\n\n",
			j,
			entry.Name(),
			info.Size(),
			info.Mode(),
			info.ModTime(),
			info.IsDir(),
			info.Sys(),
			j,
		)

		_, err = file.WriteString(textMessage)
		if err != nil {
			fmt.Printf("%q\n", err)
			return
		}
		j++
	}

	info, err := os.Stat("text1.txt")
	if err != nil {
		fmt.Printf("%q\n", err)
	} else {
		err := os.Remove(info.Name())
		if err != nil {
			fmt.Printf("%q\n", err)
		} else {
			fmt.Println("It Removed Successfully")
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cwd) // Current working directory

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(home) // User home directory

	config, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config) // OS-appropriate config root

	cache, err := os.UserCacheDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cache) // OS-appropriate cache root

	fmt.Println("Press Any Key And Enter To Quit...\n")
	fmt.Scanln()
}
