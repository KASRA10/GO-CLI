package main

import (
	"fmt"
	"os"
)

func main() {
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
}

/*
Environment and program data
Environment variables are useful for configuration such as API URLs, development modes, and secrets. Prefer LookupEnv over Getenv when an empty value has different meaning from a missing variable.

port, exists := os.LookupEnv("PORT")
if !exists {
	port = "8080"
}

fmt.Println("Server port:", port)


err := os.Setenv("APP_MODE", "development")
if err != nil {
	return
}
defer os.Unsetenv("APP_MODE")
*/

/*
Temporary files and directories
Use temporary paths for tests, uploads, conversions, and intermediate work. You must clean them up yourself.


dir, err := os.MkdirTemp("", "myapp-*")
if err != nil {
	return
}
defer os.RemoveAll(dir)

fmt.Println("Temporary directory:", dir)


file, err := os.CreateTemp("", "upload-*.txt")
if err != nil {
	return
}
defer os.Remove(file.Name())
defer file.Close()

_, err = file.WriteString("temporary content")
Temporary files are created with more restrictive permissions than ordinary created files, which is useful when their contents may be sensitive.


*/
