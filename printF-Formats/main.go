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
	fmt.Println("---------------------------------------------------")
	// String Formats
	name := "Kasra"

	fmt.Printf("%v\n", name)   // default format
	fmt.Printf("%#v\n", name)  // Go-syntax representation of the value
	fmt.Printf("%v%%\n", name) // percent sign
	fmt.Printf("%T\n", name)   // type of the value
	fmt.Printf("%s\n", name)   // string representation
	fmt.Printf("%q\n", name)   // double-quoted string representation
	fmt.Printf("%8s\n", name)  // string representation with minimum width of 8
	fmt.Printf("%-8s\n", name) // string representation with minimum width of 8 and left-justified
	fmt.Printf("%x\n", name)   // hexadecimal representation of the string
	fmt.Printf("% x\n", name)  // hexadecimal representation of the string with spaces between bytes
	/*
		Kasra
		"Kasra"
		Kasra%
		string
		Kasra
		"Kasra"
		   Kasra
		Kasra
		4b61737261
		4b 61 73 72 61
	*/
	fmt.Println("---------------------------------------------------")
	// Float Formats
	myFloat := 3.140

	fmt.Printf("%v\n", myFloat)    // default format
	fmt.Printf("%#v\n", myFloat)   // Go-syntax representation of the value
	fmt.Printf("%v%%\n", myFloat)  // percent sign
	fmt.Printf("%T\n", myFloat)    // type of the value
	fmt.Printf("%e\n", myFloat)    // scientific notation with exponent
	fmt.Printf("%E\n", myFloat)    // scientific notation with exponent (uppercase)
	fmt.Printf("%f\n", myFloat)    // decimal point but no exponent
	fmt.Printf("%.2f\n", myFloat)  // decimal point but no exponent with 2 decimal places
	fmt.Printf("%6.2f\n", myFloat) // decimal point but no exponent with 2 decimal places and minimum width of 6
	fmt.Printf("%g\n", myFloat)    // compact representation of the float with no exponent if possible
	/*
	   3.14
	   3.14
	   3.14%
	   float64
	   3.140000e+00
	   3.140000E+00
	   3.140000
	   3.14
	     3.14
	   3.14
	*/

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
