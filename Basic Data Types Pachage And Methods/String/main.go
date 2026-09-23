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
