package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		project  string
		timeout  time.Duration
		verbose  bool
		price    float64
		fileSize int64
		port     int
	)

	flag.StringVar(
		&project,
		"project",
		"flags",
		"Run Flag Package Commands",
	)
	flag.DurationVar(
		&timeout,
		"timeout",
		30*time.Second,
		"Request Timeout",
	)
	flag.BoolVar(
		&verbose,
		"verbose",
		false,
		"Verbose All Directories",
	)
	flag.Float64Var(
		&price,
		"price",
		10.97,
		"License Price",
	)
	flag.Int64Var(
		&fileSize,
		"fileSize",
		58,
		"Final Portable File Size",
	)
	flag.IntVar(
		&port,
		"port",
		90005,
		"Needed Port To Run Software",
	)
	version := flag.Float64(
		"version",
		2.1,
		"Last Updated Version Number",
	)

	flag.Parse()

	args := flag.Args()
	fmt.Println(args)

	fmt.Println(flag.NArg())
	if flag.NArg() == 0 {
		fmt.Println("No files provided")
	}

	fmt.Println(flag.NFlag())
	if flag.NFlag() == 0 {
		fmt.Println("No flags provided")
	}

	if flag.Parsed() {
		fmt.Println("Flags have been parsed")
	}

	flag.Set("version", "2.2") //* This is useful when configuration comes from somewhere other than the command line.

	flag.PrintDefaults()

	fmt.Println(*version)

	fLookUp := flag.Lookup("timeout")
	if fLookUp != nil {
		text := fmt.Sprintf(
			"Name: %q\n"+
				"Usage: %s\n"+
				"Value: %v\n"+
				"Default Value: %v\n",
			fLookUp.Name,
			fLookUp.Usage,
			fLookUp.Value,
			fLookUp.DefValue,
		)

		fmt.Println(text)
	}

	//! Visit() lets you iterate over flags that were actually set.
	flag.Visit(func(f *flag.Flag) {
		fmt.Println("Flag:", f.Name)
	})
	//* Can use -h to show defaults as well. go run main.go -h
}
