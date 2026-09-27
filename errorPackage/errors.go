package main

import (
	"errors"
	"fmt"
	"os"
)

func LoadConfig() error {
	return errors.New("Configuration File Is Missing!")
}

//* If you want a reusable error that can be compared, create a sentinel error:
//! Even if two errors have the same message, they are different values
var (
	ErrUserDoesNotExist = errors.New("User Not Found")
	ErrUserIsDisabled   = errors.New("User Is Inactive")
	ErrUserUnauthorized = errors.New("Unauthorized")
)

func FindUser(id int) error {
	if id < 10 {
		return ErrUserUnauthorized
	} else if id > 10 && id < 20 {
		return ErrUserIsDisabled
	} else if id > 20 {
		return ErrUserDoesNotExist
	} else {
		return errors.New("Service Not Available")
	}
}

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
)

func ValidateForm() error {
	var errs []error

	errs = append(errs, ErrInvalidUsername)
	errs = append(errs, ErrInvalidPassword)

	return errors.Join(errs...)
}

func main() {
	/*
		Creating errors with errors.New
		Syntax
		go
		func New(text string) error
		errors.New creates an error containing the specified text.
	*/

	err := errors.New("This Is First Error")

	fmt.Println(err)

	err = LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(FindUser(100))
	fmt.Println(FindUser(1))
	fmt.Println(FindUser(10))

	userName := "Kasra10"
	id := 10

	err = fmt.Errorf("User %q with ID: %d was Not Found", userName, id)
	fmt.Println(err)

	/*
		Wrapping errors with %w
		Wrapping means adding context while preserving the original error.
	*/

	ErrDataBase := errors.New("DataBase Connection Failed")

	fmt.Println(fmt.Errorf("InterNet Connection: %w", ErrDataBase))

	/*
		Checking errors with errors.Is
		func Is(err, target error) bool
		errors.Is checks whether an error or any error it wraps matches a target error.
		It should generally be preferred over direct equality because it works through wrapped errors.
	*/

	error1 := errors.New("Original Error")
	error2 := fmt.Errorf("Second Layer: %w", error1)
	error3 := fmt.Errorf("Third Layer: %w", error2)

	fmt.Println(errors.Is(error3, error2)) // true
	fmt.Println(errors.Is(error2, error1)) // true
	fmt.Println(errors.Is(error1, error3)) // false

	_, err = os.Open("missing.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("The File Does Not Exist In This Directory")
	}
	fmt.Println("Other Error:", err)

	/*
		Combining errors with errors.Join
		Syntax

		func Join(errs ...error) error
		errors.Join combines multiple errors into one error.
		It was added in Go 1.20. Nil errors are ignored, and Join returns nil if all supplied errors are nil.
	*/

	err = ValidateForm()

	fmt.Println(err)
	/*
		it print both:
		invalid username
		invalid password
	*/
	if errors.Is(err, ErrInvalidUsername) {
		fmt.Println("Username is invalid")
	}

	if errors.Is(err, ErrInvalidPassword) {
		fmt.Println("Password is invalid")
	}

	/*
		Unwrapping errors with errors.Unwrap
		Syntax
		func Unwrap(err error) error
		errors.Unwrap returns the error wrapped by an error implementing:

		Unwrap() error
	*/

	error1 = errors.New("Original Error")
	error2 = fmt.Errorf("Second Layer: %w", error1)
	error3 = fmt.Errorf("Third Layer: %w", error2)

	if errors.Is(error2, error1) {
		fmt.Println("true, Is wrapped")
		fmt.Println("UnWrappedPart:", errors.Unwrap(error2))
	}
	//? Can Use FOr Joined Errors AS well
}

/*
Go errors Package
The standard-library errors package provides functions for
creating, wrapping, comparing, extracting, and combining errors.
In the current Go documentation, its public API contains the variable ErrUnsupported and six functions:
New, Is, As, AsType, Join, and Unwrap
*/

/*
Important: every call creates a distinct error
Even if two errors have the same message, they are different values:

err1 := errors.New("not found")
err2 := errors.New("not found")

fmt.Println(err1 == err2)        // false
fmt.Println(errors.Is(err1, err2)) // false
*/

/*
fmt.Errorf()

| Verb | Meaning                |
| ---- | ---------------------- |
| %s   | String                 |
| %d   | Integer                |
| %f   | Floating-point number  |
| %v   | Default representation |
| %q   | Quoted representation  |
| %T   | Type of a value        |
| %w   | Wrap another error     |
*/

/*
 Extracting typed errors with errors.As
Syntax
go
func As(err error, target any) bool
errors.As searches an error chain for a particular error type. If it finds one, it assigns the matching error to target.

A common example is *fs.PathError:

go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	_, err := os.Open("missing.txt")

	if err != nil {
		var pathErr *fs.PathError

		if errors.As(err, &pathErr) {
			fmt.Println("Operation:", pathErr.Op)
			fmt.Println("Path:", pathErr.Path)
			fmt.Println("Underlying error:", pathErr.Err)
		}
	}
}
Possible output:

text
Operation: open
Path: missing.txt
Underlying error: file does not exist
Why use As?
Use errors.Is when you want to ask:

Is this error a particular known condition?

Use errors.As when you want to ask:

Is this error a particular type, and can I access its fields?

For example:

if errors.Is(err, fs.ErrNotExist) {
	// Check a condition.
}

var pathErr *fs.PathError
if errors.As(err, &pathErr) {
	// Access pathErr.Path, pathErr.Op, and pathErr.Err.
}
*/

/*
Generic extraction with errors.AsType
Syntax
go
func AsType[E error](err error) (E, bool)
errors.AsType was added in Go 1.26. It searches the error tree for a matching error type and returns the error directly, without requiring a target pointer.

Instead of writing:

go
var pathErr *fs.PathError

if errors.As(err, &pathErr) {
	fmt.Println(pathErr.Path)
}
You can write:

go
if pathErr, ok := errors.AsType[*fs.PathError](err); ok {
	fmt.Println(pathErr.Path)
}
Complete example:

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	_, err := os.Open("missing.txt")

	if err != nil {
		pathErr, ok := errors.AsType[*fs.PathError](err)

		if ok {
			fmt.Println("Failed operation:", pathErr.Op)
			fmt.Println("Failed path:", pathErr.Path)
		}
	}
}
As versus AsType
Feature	errors.As	errors.AsType
Available since	Go 1.13	Go 1.26
Style	Uses a target pointer	Uses a generic type
Example	errors.As(err, &target)	errors.AsType[*MyError](err)
Readability	More verbose	Usually simpler
Compatibility	Works with older Go versions	Requires Go 1.26 or newer
If you need your project to support older Go versions, use errors.As.
*/
