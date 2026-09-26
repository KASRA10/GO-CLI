package main

import (
	"fmt"
	"os"
)

func main() {
	// Printf is used to format and print values to the standard output.
	// It allows you to specify a format string that defines how the values should be formatted.
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

	// Build a string: Sprint, Sprintln, Sprintf
	// Use these when you need the formatted result in a variable instead of printing it.
	name = "kasra"
	score := 100

	message := fmt.Sprint("Welcome ", name)
	line := fmt.Sprint("Score Is: ", score)

	fmt.Println(message)
	fmt.Println(line)

	/*
		* Sprintln includes a trailing newline. Sprintf is especially useful for
		* HTML templates, log messages, file names, SQL placeholders in examples, CLI display,
		* and API messages
		! but do not use it to concatenate untrusted data into SQL queries.
	*/

	file, err := os.Create("report.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "This is a sample report.")
	fmt.Fprintf(file, "Total sales: %.2f\n", 1500.75)

	data, err := os.ReadFile("report.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Println("Report content:")
	fmt.Println(string(data))

	var userInput string

	fmt.Print("Enter A Word: ")

	_, err = fmt.Scanln(&userInput)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	var age int

	fmt.Print("Enter your age: ")

	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("Please enter a whole number.")
		return
	}

	fmt.Println("Next year:", age+1)
	/*
		* Scan treats newlines as spaces.
		* * Scanln reads until a newline.
		* * * Scanf follows a format pattern.
		! These are okay for simple CLI practice;
		! for realistic user input with spaces and validation, use bufio.Reader plus strings and strconv.
	*/

	//? Sscan... reads values from a string instead of the terminal.
	input := "Kasra 24 19.99"

	var price float64

	count, err := fmt.Sscan(input, &name, &age, &price)

	fmt.Println(count, err)       // 3 <nil>
	fmt.Println(name, age, price) // Kasra 24 19.99

	input = "id=42"

	var id int

	count, err = fmt.Sscanf(input, "id=%d", &id)

	fmt.Println(count, err, id) // 1 <nil> 42

	/*
		* Append to a byte slice
		? Append, Appendln, and Appendf add formatted output to an existing []byte.
		? They are useful when building buffers efficiently.
	*/
	buffer := []byte("User: ")

	buffer = fmt.Append(buffer, "Kasra")
	buffer = fmt.Appendln(buffer, "| Status:", "active")
	buffer = fmt.Appendf(buffer, "Score: %d", 95)

	fmt.Println(string(buffer))
	// User: Kasra| Status: active
	// Score: 95

	data, err = os.ReadFile("report.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Println("Report content:")
	fmt.Println(string(data))

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}

/*
Rarely used: FormatString
fmt.FormatString is mainly for people implementing custom Go formatters with the fmt.Formatter interface. It returns the active formatting directive from the formatting state and verb, such as %+v or %.2f. Most application code does not need it.

Daily-use priority list
For your Go CLI, backend, and file exercises, master these first:

fmt.Println — quick output and debugging.

fmt.Printf — terminal output with verbs.

fmt.Sprintf — create formatted strings.

fmt.Errorf with %w — meaningful, wrappable errors.

fmt.Fprintln and fmt.Fprintf — write formatted content to files and HTTP responses.

fmt.Scanln or fmt.Sscanf — simple learning-level input parsing.

Given this struct, how would you print the field names and values in one line using fmt.Printf and the %+v verb?

go
user := struct {
	Name string
	Age  int
}{
	Name: "Kasra",
	Age:  24,
}
*/
