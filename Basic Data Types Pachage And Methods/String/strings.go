package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	myString := "Hello GO Developers!"

	// Clone: independent copy of text
	clonedText := strings.Clone(myString)
	fmt.Println(clonedText)

	// Compare: 0 = equal, -1 = first comes before second, +1 = first comes after
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("c", "b"))
	fmt.Println(strings.Compare("Go", "go"))
	fmt.Println(strings.Compare("Go", "Go"))
	/*
		-1
		1
		-1
		0
	*/

	// Contains
	fmt.Println(strings.Contains(myString, "World")) // false

	// ContainsAny: contains at least one character from the second string
	fmt.Println(strings.ContainsAny(myString, "DGo")) // true

	// ContainsRune: rune is a Unicode character
	fmt.Println(strings.ContainsRune(myString, 'G')) // true

	// ContainsFunc: check whether at least one rune matches a rule
	fmt.Println(strings.ContainsFunc(myString, unicode.IsDigit)) // false

	// Count occurrences
	fmt.Println(strings.Count("Go, This Is Go For Go", "Go")) // 3

	// Start/end checks
	fmt.Println(strings.HasPrefix(myString, "Hello"))      // true
	fmt.Println(strings.HasSuffix(myString, "Developers")) // false

	// Case-insensitive equality
	fmt.Println(strings.EqualFold("GO", "go")) // true

	// First occurrence
	fmt.Println(strings.Index(myString, "GO"))     // 6
	fmt.Println(strings.Index(myString, "Google")) // -1

	// Last occurrence
	fmt.Println(strings.LastIndex("GO With GO", "GO")) // 8

	// First occurrence of any listed character
	fmt.Println(strings.IndexAny("Kasra", "xyzs")) // 2

	// Last occurrence of any listed character
	fmt.Println(strings.LastIndexAny("go.dev", "./")) // 2

	// Search for one byte: good for ASCII characters
	fmt.Println(strings.IndexByte(myString, '1'))     // -1
	fmt.Println(strings.LastIndexByte(myString, '1')) // -1

	// Search for one rune: appropriate for Unicode characters too
	fmt.Println(strings.IndexRune("سلام دنیا!", 'س')) // 0

	// Find first digit using a rule
	fmt.Println(strings.IndexFunc("Product-110", unicode.IsDigit)) // 8

	// Find last digit using a rule
	fmt.Println(strings.LastIndexFunc("Product-110", unicode.IsDigit)) // 10

	text := "GO=Developers=GODEV"
	// Cut: divide at the FIRST "="
	key, value, found := strings.Cut(text, "=")
	fmt.Println(key, value, found) // GO Developers=GODEV true

	before, after, found := strings.CutLast(text, "=")
	fmt.Println(before, after, found) // GO=Developers GODEV true

	// CutPrefix: remove prefix and report success
	beforePrefix, found := strings.CutPrefix(text, "GO=Developers=")
	fmt.Println(beforePrefix, found) // GODEV true

	// CutSuffix: remove suffix and report success
	beforeSuffix, found := strings.CutSuffix(text, "=GODEV")
	fmt.Println(beforeSuffix, found) // GO=Developers trues

	// Split: all pieces
	split := strings.Split("go, php, mysql", ",")
	fmt.Println(split)      // [go  php  mysql]
	fmt.Printf("%T", split) // []string

	// SplitN: at most n pieces
	fmt.Println(strings.SplitN("go, php, js, mysql", ",", 2)) // []string[go  php, js, mysql]

	// SplitAfterN: separator stays attached, at most n pieces
	fmt.Println(strings.SplitAfterN("go, php, js, mysql", ",", 2)) // [go,  php, js, mysql]

	// Fields: split by one or more whitespace characters
	fmt.Println(strings.Fields("   This\nIs  Go\tDeveloping ")) // [This Is Go Developing]

	// Join: reverse of Split
	languages := []string{"GO", "PHP", "MySQL"}
	fmt.Printf("%q\n", strings.Join(languages, "|")) // "GO|PHP|MySQL"

	text = "GO Is Great. GO Is Fast!"

	// Replace every match
	fmt.Println(strings.ReplaceAll(text, "GO", "GoLang")) // GoLang Is Great. GoLang Is Fast!

	// Replace only the first match
	fmt.Println(strings.Replace(text, "GO", "GoLang", 1)) // GoLang Is Great. GO Is Fast!

	// Repeat
	fmt.Println(strings.Repeat("*", 25)) // *************************

	// Lowercase / uppercase
	fmt.Println(strings.ToLower(text)) // go is great. go is fast!
	fmt.Println(strings.ToUpper(text)) // GO IS GREAT. GO IS FAST!

	// ToTitle changes every letter to title case
	fmt.Println(strings.ToTitle(text)) // GO IS GREAT. GO IS FAST!

	// NewReplacer: replace several pairs efficiently
	replacer := strings.NewReplacer(
		"Go", "Golang",
		"PHP", "PHP 8",
		"JS", "JavaScript",
	)

	fmt.Println(replacer.Replace("Go, PHP, JS"))
	// Golang, PHP 8, JavaScript

	text = "  !! hello Go !!  "

	// Remove outside whitespace
	fmt.Println(strings.TrimSpace(text)) // !! hello Go !!

	// Remove spaces and ! from BOTH ends
	fmt.Println(strings.Trim(text, " !")) // hello Go

	// Only left side
	fmt.Println(strings.TrimLeft("---title", "-")) // title

	// Only right side
	fmt.Println(strings.TrimRight("title---", "-")) // title

	// Exact prefix and suffix
	fmt.Println(strings.TrimPrefix("https://example.com", "https://"))
	// example.com

	fmt.Println(strings.TrimSuffix("photo.png", ".png"))
	// photo

	// Trim according to a rule: remove starting/ending digits
	fmt.Println(strings.TrimFunc("123hello456", unicode.IsDigit))
	// hello

	fmt.Println(strings.TrimLeftFunc("123hello", unicode.IsDigit))
	// hello

	fmt.Println(strings.TrimRightFunc("hello456", unicode.IsDigit))
	// hello

	// NewReader: makes a string act like an io.Reader.
	// Helpful when another function expects a reader.
	reader := strings.NewReader("Hello Go")

	buffer := make([]byte, 5)
	reader.Read(buffer)

	fmt.Println(string(buffer)) // Hello

	// ToValidUTF8: replaces invalid UTF-8 bytes.
	broken := string([]byte{'G', 'o', 0xff, '!'})
	clean := strings.ToValidUTF8(broken, "?")

	fmt.Println(clean) // Go?
}

/*
1. strings (String Manipulation)
The strings package implements simple functions to manipulate UTF‑8 encoded strings.

Functions
Function	Signature	Added
Clone	func Clone(s string) string	go1.18
Compare	func Compare(a, b string) int	go1.5
Contains	func Contains(s, substr string) bool	–
ContainsAny	func ContainsAny(s, chars string) bool	–
ContainsFunc	func ContainsFunc(s string, f func(rune) bool) bool	go1.21
ContainsRune	func ContainsRune(s string, r rune) bool	–
Count	func Count(s, substr string) int	–
Cut	func Cut(s, sep string) (before, after string, found bool)	go1.18
CutLast	func CutLast(s, sep string) (before, after string, found bool)	go1.27
CutPrefix	func CutPrefix(s, prefix string) (after string, found bool)	go1.20
CutSuffix	func CutSuffix(s, suffix string) (before string, found bool)	go1.20
EqualFold	func EqualFold(s, t string) bool	–
Fields	func Fields(s string) []string	–
FieldsFunc	func FieldsFunc(s string, f func(rune) bool) []string	–
HasPrefix	func HasPrefix(s, prefix string) bool	–
HasSuffix	func HasSuffix(s, suffix string) bool	–
Index	func Index(s, substr string) int	–
IndexAny	func IndexAny(s, chars string) int	–
IndexByte	func IndexByte(s string, c byte) int	–
IndexFunc	func IndexFunc(s string, f func(rune) bool) int	–
IndexRune	func IndexRune(s string, r rune) int	–
Join	func Join(elems []string, sep string) string	–
LastIndex	func LastIndex(s, substr string) int	–
LastIndexAny	func LastIndexAny(s, chars string) int	–
LastIndexByte	func LastIndexByte(s string, c byte) int	–
LastIndexFunc	func LastIndexFunc(s string, f func(rune) bool) int	–
Map	func Map(f func(rune) rune, s string) string	–
NewReader	func NewReader(s string) *Reader	–
NewReplacer	func NewReplacer(oldnew ...string) *Replacer	–
Repeat	func Repeat(s string, count int) string	–
Replace	func Replace(s, old, new string, n int) string	–
ReplaceAll	func ReplaceAll(s, old, new string) string	go1.12
Split	func Split(s, sep string) []string	–
SplitAfter	func SplitAfter(s, sep string) []string	–
SplitAfterN	func SplitAfterN(s, sep string, n int) []string	–
SplitN	func SplitN(s, sep string, n int) []string	–
ToLower	func ToLower(s string) string	–
ToLowerSpecial	func ToLowerSpecial(c unicode.SpecialCase, s string) string	–
ToTitle	func ToTitle(s string) string	–
ToTitleSpecial	func ToTitleSpecial(c unicode.SpecialCase, s string) string	–
ToUpper	func ToUpper(s string) string	–
ToUpperSpecial	func ToUpperSpecial(c unicode.SpecialCase, s string) string	–
ToValidUTF8	func ToValidUTF8(s, replacement string) string	–
Trim	func Trim(s, cutset string) string	–
TrimFunc	func TrimFunc(s string, f func(rune) bool) string	–
TrimLeft	func TrimLeft(s, cutset string) string	–
TrimLeftFunc	func TrimLeftFunc(s string, f func(rune) bool) string	–
TrimPrefix	func TrimPrefix(s, prefix string) string	–
TrimRight	func TrimRight(s, cutset string) string	–
TrimRightFunc	func TrimRightFunc(s string, f func(rune) bool) string	–
TrimSpace	func TrimSpace(s string) string	–
TrimSuffix	func TrimSuffix(s, suffix string) string
*/
