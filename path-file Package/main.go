package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	/*
		What path/filepath is for
		path/filepath gives you OS-aware utilities to build, split, clean, and match file system paths.
		It automatically uses / on Unix-like systems and \ on Windows,
		and knows about things like drive letters (C:) and UNC paths (\\host\share).
		Use it instead of manually concatenating strings with / or \.
	*/

	// newPath := filepath.Join("workspace", "data", "settings.json")
	// fmt.Println(newPath) // workspace\data\settings.json

	newPath := filepath.Join(`D:\`, "Kasra10", "Others") // D:\Kasra10\Others

	entires, err := os.ReadDir(newPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entires {
		fmt.Println(entry.Name())
	}
	//! if the directory does not exist: Error reading directory: open D:Kasra10\Others: The system cannot find the path specified.

	/*
		filepath.Base
		Base gives you the last element of a path.
	*/

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cwd)

	lastElement := filepath.Base(cwd)
	fmt.Println(lastElement)

	/*
		filepath.Dir
		Dir gives you the directory portion.
	*/
	dirPortion := filepath.Dir(cwd)
	fmt.Println(dirPortion)

	/*
		filepath.Ext
		Ext gets the file extension.
	*/

	file, err := os.Create("setting.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Close()

	fmt.Println(filepath.Ext(cwd + `\setting.json`)) // .json

	/*
		filepath.Split
		Split gives you directory + filename.
	*/

	myPath := cwd + `\setting.json`

	dir, ext := filepath.Split(myPath)
	fmt.Println("Directory:", dir)
	fmt.Println("File:", ext)
	/*
		Directory: D:\Kasra10\Others\K10Git\GO-CLI\path-file Package\
		File: setting.json
	*/

	/*
		*Dir
		?Give me the parent directory.
		*Split
		?Separate this path into directory and filename.
	*/

	/*
		filepath.Clean
		Clean normalizes a path.
	*/

	path := "workspace//data/./files/../settings.json"

	fmt.Println(filepath.Clean(path)) // workspace\data\settings.json

	/*
		filepath.Abs
		Abs converts a path into an absolute path.
	*/
	absolutePath, err := filepath.Abs("workspace/data/settings.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("absolutePath:\n", absolutePath) // D:\Kasra10\Others\K10Git\GO-CLI\path-file Package\workspace\data\settings.json

	/*
		filepath.IsAbs
		Checks whether a path is absolute.
		It's simply a boolean test.
	*/

	fmt.Println(filepath.IsAbs("workspace/data/file.txt")) // false
	fmt.Println(filepath.IsAbs("/home/user/file.txt"))     // false
	fmt.Println(filepath.IsAbs(cwd))                       // true

	/*
		filepath.Rel
		Rel calculates a relative path between two paths.
	*/

	base := "/home/user/project"
	target := "/home/user/project/data/settings.json"

	relative, err := filepath.Rel(base, target)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(relative) // data\settings.json

	/*
		filepath.ToSlash
		Converts OS-specific path separators to /.
	*/

	path = `C:\Users\Kasra\Documents\file.txt`

	fmt.Println(filepath.ToSlash(path)) // C:/Users/Kasra/Documents/file.txt

	/*
		filepath.FromSlash
		This does the opposite.
	*/

	path = "workspace/data/settings.json"

	fmt.Println(filepath.FromSlash(path)) // workspace\data\settings.json

	/*
		filepath.VolumeName
		This is particularly useful for Windows.
	*/

	path = `C:\Users\Kasra\file.txt`

	fmt.Println(filepath.VolumeName(path)) // C:

	/*
		filepath.Match
		Match checks whether a filename/path matches a shell-style pattern.
	*/

	matched, err := filepath.Match("*.txt", "hello.txt")

	fmt.Println(matched) // true
	fmt.Println(err)     // nil

	//? ? Matches exactly one character.
	filepath.Match("file?.txt", "file1.txt") // true

	filepath.Match("file[0-9].txt", "file5.txt") // true

	/*
		filepath.Glob
		Glob is like Match, but it actually searches filesystem paths.
	*/
	matches, err := filepath.Glob("workspace/data/*.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range matches {
		fmt.Println(file)
	}
	// nil

	/*
		Difference
		? Match:
		* Does this name match this pattern?

		? Glob:
		* Find files that match this pattern.

		* Glob uses the same pattern syntax as Match.
		! If nothing matches, it returns nil without treating that as an error.
		! The documented possible error is ErrBadPattern for malformed patterns
	*/

	/*
		filepath.IsLocal
		This is a newer and very interesting function.
		It checks whether a path is local and relative, using lexical analysis.
		* The official definition includes that the path must not be absolute, must not be empty,
		* and must stay within the subtree represented by the path.
		* On Windows it also rejects reserved names such as NUL.
	*/

	fmt.Println(filepath.IsLocal("data/file.txt")) // true
	fmt.Println(filepath.IsLocal("../file.txt"))   // false
	fmt.Println(filepath.IsLocal("/etc/passwd"))   // false

	/**/

	userPath := "documents/report.pdf"

	if !filepath.IsLocal(userPath) {
		fmt.Println("Invalid path")
		return
	}

	fullPath := filepath.Join("uploads", userPath)

	fmt.Println(fullPath) // uploads\documents\report.pdf
	//* This is particularly relevant to security-sensitive file handling.

	/*
		filepath.Localize
		* It converts a slash-separated path into an OS-specific filesystem path.
	*/

	path, err = filepath.Localize("workspace/data/settings.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(path) // workspace\data\settings.json

	/*
		* filepath.EvalSymlinks
		* This one is more advanced.
		* A symbolic link is essentially a filesystem object that points to another file/directory.
		* The function evaluates symbolic links in the path and returns the resulting path. It also cleans the result.
		* his is useful when you need to know:
		? "What filesystem object does this path ultimately point to?"
	*/

	realPath, err := filepath.EvalSymlinks("data/current.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(realPath) // GetFileAttributesEx data: The system cannot find the filespecified.

	/*
		filepath.SkipDir
		You can tell Walk:
		Don't enter this directory.
	*/

	// err = filepath.Walk("workspace", func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if info.IsDir() && info.Name() == "node_modules" {
	// 		return filepath.SkipDir
	// 	}

	// 	fmt.Println(path)

	// 	return nil
	// })

	/*
		30. filepath.ListSeparator
		This is the separator between multiple paths.
		Unix:
		:

		Windows:
		;

		For example, Windows PATH might look like:
		C:\Go\bin;C:\Windows\System32

		Unix:
		/usr/local/bin:/usr/bin

		This is why SplitList knows how to correctly split environment-style path lists. Go Packages
	*/
}
