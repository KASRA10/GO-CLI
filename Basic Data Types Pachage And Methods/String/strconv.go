package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string -> int
	age, err := strconv.Atoi("24")
	fmt.Println(age, err) // 24 <nil>

	age, err = strconv.Atoi("String")
	fmt.Println(age, err) // 0 strconv.Atoi: parsing "String": invalid syntax

	// int -> string
	textAge := strconv.Itoa(31)
	fmt.Println(textAge) // 31

	// string -> float64
	price, err := strconv.ParseFloat("99.99", 64)
	fmt.Println(price, err) // 99.99 <nil>

	// float64 -> string
	textPrice := strconv.FormatFloat(99.99999, 'f', 2, 64) // 100.00
	fmt.Println(textPrice)

	textPrice = strconv.FormatFloat(99.93, 'f', 1, 64) // 99.9
	fmt.Println(textPrice)

	// string -> bool
	enabled, err := strconv.ParseBool("true")
	fmt.Println(enabled, err) // true <nil>

	// bool -> string
	textEnabled := strconv.FormatBool(true)
	fmt.Println(textEnabled) // "true"

	/*
		Parse... functions can fail,
		so always check err before trusting user input, form values, query parameters, environment variables,
		or file content.
	*/

	// "1010" in base 2 -> decimal 10
	binaryValue, err := strconv.ParseInt("1010", 2, 64) // number, base, intBASE
	fmt.Println(binaryValue, err)                       // 10 <nil>

	// "ff" in base 16 -> decimal 255
	hexValue, err := strconv.ParseInt("fafafa", 16, 64)
	fmt.Println(hexValue, err) // 16448250 <nil>

	// decimal 255 -> hexadecimal text
	hexText := strconv.FormatInt(255, 16)
	fmt.Println(hexText) // "ff"

	// Unsigned integer: cannot be negative
	id, err := strconv.ParseUint("500", 10, 64)
	fmt.Println(id, err) // 500 <nil>

	fmt.Println(strconv.FormatUint(500, 2)) // "111110100"

	// strconv.FormatFloat(number, format, precision, bitSize)
	n := 1234.56789

	fmt.Println(strconv.FormatFloat(n, 'f', 2, 64))  // 1234.57
	fmt.Println(strconv.FormatFloat(n, 'f', 0, 64))  // 1235
	fmt.Println(strconv.FormatFloat(n, 'e', 2, 64))  // 1.23e+03
	fmt.Println(strconv.FormatFloat(n, 'E', 2, 64))  // 1.23E+03
	fmt.Println(strconv.FormatFloat(n, 'g', -1, 64)) // 1234.56789
	/*
		Use 'f' for UI prices and ordinary decimal numbers.
		Use -1 precision when you want a compact representation that preserves the value
		without choosing a fixed number of decimal places.
	*/

	value, err := strconv.ParseBool("TRUE")
	fmt.Println(value, err) // true <nil>

	value, err = strconv.ParseBool("yes")
	fmt.Println(value, err) // false + error

	number, err := strconv.ParseComplex("2.5+3i", 128)
	fmt.Println(number, err) // (2.5+3i) <nil>

	text := strconv.FormatComplex(2.5+3i, 'f', 1, 128)
	fmt.Println(text) // "(2.5+3.0i)"

	text = "Hello\nKasra"

	quoted := strconv.Quote(text)
	fmt.Println(quoted) // "Hello\nKasra"

	original, err := strconv.Unquote(quoted)
	fmt.Println(original, err)
	// Hello
	// Kasra <nil>

	fmt.Println(strconv.QuoteRune('G'))       // 'G'
	fmt.Println(strconv.QuoteRune('س'))       // 'س'
	fmt.Println(strconv.QuoteToASCII("سلام")) // "\u0633\u0644\u0627\u0645"

	/*
		Append functions
		The Append... family appends a converted value to a []byte.
		This is useful for performance-sensitive code because
		it can reuse an existing byte buffer instead of creating many separate strings.
	*/

	buf := []byte("User: ")

	buf = strconv.AppendInt(buf, 31, 10)
	buf = append(buf, ' ')

	buf = strconv.AppendBool(buf, true)
	buf = append(buf, ' ')

	buf = strconv.AppendFloat(buf, 20.35, 'f', 2, 64)
	buf = append(buf, ' ')

	fmt.Println(string(buf)) // User: 31 true 20.35

	fmt.Println(strconv.IsPrint('A'))  // true
	fmt.Println(strconv.IsPrint('\n')) // false

	fmt.Println(strconv.IsGraphic('س')) // true
	fmt.Println(strconv.IsGraphic(' ')) // true
	/*
	   IsPrint checks whether a rune is printable, excluding control characters;
	   IsGraphic includes Unicode graphic characters and spaces.
	*/
}
