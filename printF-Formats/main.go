package main

import (
	"fmt"
)

func main() {
	// INT Formats
	myINT := 77
	var myHex int = 0xfafafa

	fmt.Printf("%v\n", myINT)   // default format
	fmt.Printf("%#v\n", myINT)  // Go-syntax representation of the value
	fmt.Printf("%v%%\n", myINT) // percent sign
	fmt.Printf("%T\n", myINT)   // type of the value
	fmt.Printf("%b\n", myINT)   // binary representation
	fmt.Printf("%d\n", myINT)   // decimal representation
	fmt.Printf("%d\n", myHex)   // decimal representation
	fmt.Printf("%+d\n", myINT)  // Signeddecimal representation
	fmt.Printf("%o\n", myINT)   // octal representation
	fmt.Printf("%O\n", myINT)   // octal representation with leading 0o
	fmt.Printf("%x\n", myINT)   // hexadecimal representation
	fmt.Printf("%X\n", myINT)   // hexadecimal UpperCase representation
	fmt.Printf("%#X\n", myINT)  // hexadecimal UpperCase representation with leading 0x
	fmt.Printf("%4d\n", myINT)  // decimal representation with minimum width of 4
	fmt.Printf("%-4d\n", myINT) // decimal representation with minimum width of 4 and left-justified
	fmt.Printf("%04d\n", myINT) // decimal representation with minimum width of 4 and zero-padded
	/*
		77
		77
		77%
		int
		1001101
		77
		16448250
		+77
		115
		0o115
		4d
		4D
		0X4D
		  77
		77
		0077
	*/
}
