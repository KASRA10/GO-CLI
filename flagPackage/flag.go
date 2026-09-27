package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String(
		"name",
		"Kasra",
		"NewName",
	)

	flag.Parse()

	fmt.Println("Hello", *name) //* go run flag.go -name=David ==> Hello David
	// fmt.Println("Hello", *Kasra)
	// fmt.Println("Hello", *newName)
	/*
		.\flag.go:18:24: undefined: Kasra
		.\flag.go:19:24: undefined: newName
	*/
}

/*
The flag package is one of the most useful standard-library packages for building CLI applications in Go, especially since you want to learn CLI development before moving deeper into APIs and microservices.
I checked the current official Go documentation for flag (Go 1.27.1). Go Packages
Go flag Package — Complete Guide
Package:
import "flag"

Its main purpose is:
Read options/parameters supplied to a program from the command line.

For example, suppose you build:
backup.exe

You can run:
backup.exe -source=data -destination=backup

Your Go program can then read:
source      = data
destination = backup

The flag package handles parsing these values for you. Go Packages
1. The Basic Concept
There are two important concepts:
Flag
A named option:
-name=value

Example:
-port=8080

Argument
A value that comes after the flags but isn't itself a flag.
Example:
myapp -port=8080 file.txt

Here:
-port=8080

is a flag.
file.txt

is an argument.
The flag package supports both
*/

/*
Understanding flag.String
This:
name := flag.String(
	"name",
	"Kasra",
	"Your name",
)

has three important parameters:
flag.String(name, defaultValue, usage)

Specifically:
flag.String(
	"name",     // flag name
	"Kasra",    // default value
	"Your name", // description
)

It returns:
*string

So:
fmt.Println(*name)

not:
fmt.Println(name)

because name is a pointer.
*/

/*
Absolutely. The `flag` package is one of the most useful standard-library packages for building **CLI applications in Go**, especially since you want to learn CLI development before moving deeper into APIs and microservices.

I checked the current official Go documentation for `flag` (Go 1.27.1). :chatgpt-content-reference{index="0"}

# Go `flag` Package — Complete Guide

Package:

```go
import "flag"
```

Its main purpose is:

> **Read options/parameters supplied to a program from the command line.**

For example, suppose you build:

```text
backup.exe
```

You can run:

```bash
backup.exe -source=data -destination=backup
```

Your Go program can then read:

```text
source      = data
destination = backup
```

The `flag` package handles parsing these values for you. :chatgpt-content-reference{index="1"}

---

# 1. The Basic Concept

There are two important concepts:

### Flag

A named option:

```bash
-name=value
```

Example:

```bash
-port=8080
```

### Argument

A value that comes after the flags but isn't itself a flag.

Example:

```bash
myapp -port=8080 file.txt
```

Here:

```text
-port=8080
```

is a **flag**.

```text
file.txt
```

is an **argument**.

The `flag` package supports both. :chatgpt-content-reference{index="2"}

---

# 2. Your First Flag

The simplest example:

```go
package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String(
		"name",
		"Kasra",
		"Your name",
	)

	flag.Parse()

	fmt.Println("Hello", *name)
}
```

Run:

```bash
go run main.go
```

Output:

```text
Hello Kasra
```

Run:

```bash
go run main.go -name=Ali
```

Output:

```text
Hello Ali
```

---

# 3. Understanding `flag.String`

This:

```go
name := flag.String(
	"name",
	"Kasra",
	"Your name",
)
```

has three important parameters:

```go
flag.String(name, defaultValue, usage)
```

Specifically:

```go
flag.String(
	"name",     // flag name
	"Kasra",    // default value
	"Your name", // description
)
```

It returns:

```go
*string
```

So:

```go
fmt.Println(*name)
```

not:

```go
fmt.Println(name)
```

because `name` is a pointer.

---

# 4. What Is the Default Value?

This is especially important.

When you write:

```go
name := flag.String("name", "Kasra", "Your name")
```

the second parameter:

```go
"Kasra"
```

is the **default value**.

If the user doesn't provide:

```bash
-name
```

then:

```text
name = Kasra
```

If the user does:

```bash
-name=Ali
```

then:

```text
name = Ali
```

So:

```text
flag.String("name", "Kasra", "Your name")
                    ↑
              default value
```

The official documentation describes the second parameter as the flag's default value. :chatgpt-content-reference{index="3"}

---

# 5. How Do I Get the Default Value?

There are actually several ways to think about this.

## Method 1 — You already know it

If you write:

```go
name := flag.String("name", "Kasra", "Your name")
```

then:

```go
"Kasra"
```

is the default.

You can create a constant:

```go
const defaultName = "Kasra"

name := flag.String(
	"name",
	defaultName,
	"Your name",
)
```

This is often cleaner.

---

# 6. Getting the Current Value

After:

```go
flag.Parse()
```

you can access:

```go
fmt.Println(*name)
```

Suppose:

```go
name := flag.String("name", "Kasra", "Your name")

flag.Parse()

fmt.Println(*name)
```

Running:

```bash
go run main.go
```

gives:

```text
Kasra
```

Running:

```bash
go run main.go -name=Ali
```

gives:

```text
Ali
```

So the variable contains:

> **The default value if the user didn't provide the flag, otherwise the user-provided value.**

---

# 7. `String`

Signature:

```go
flag.String(name string, value string, usage string) *string
```

Use it for text.

Example:

```go
username := flag.String(
	"username",
	"guest",
	"Username",
)

flag.Parse()

fmt.Println(*username)
```

Command:

```bash
go run main.go -username=admin
```

Output:

```text
admin
```

---

# 8. `Int`

For integers:

```go
age := flag.Int(
	"age",
	18,
	"User age",
)

flag.Parse()

fmt.Println(*age)
```

Run:

```bash
go run main.go -age=31
```

Output:

```text
31
```

Without the flag:

```bash
go run main.go
```

Output:

```text
18
```

Signature:

```go
flag.Int(name string, value int, usage string) *int
```

---

# 9. `Int64`

For `int64`:

```go
id := flag.Int64(
	"id",
	1000,
	"User ID",
)

flag.Parse()

fmt.Println(*id)
```

Example:

```bash
go run main.go -id=987654321
```

---

# 10. `Uint`

For unsigned integers:

```go
count := flag.Uint(
	"count",
	10,
	"Number of items",
)

flag.Parse()

fmt.Println(*count)
```

Example:

```bash
go run main.go -count=50
```

---

# 11. `Uint64`

For `uint64`:

```go
size := flag.Uint64(
	"size",
	1024,
	"File size",
)

flag.Parse()

fmt.Println(*size)
```

---

# 12. `Float64`

For decimal numbers:

```go
price := flag.Float64(
	"price",
	10.5,
	"Product price",
)

flag.Parse()

fmt.Println(*price)
```

Run:

```bash
go run main.go -price=25.75
```

Output:

```text
25.75
```

---

# 13. `Bool`

Boolean flags are slightly different.

```go
verbose := flag.Bool(
	"verbose",
	false,
	"Enable verbose output",
)

flag.Parse()

fmt.Println(*verbose)
```

Without the flag:

```bash
go run main.go
```

Output:

```text
false
```

With:

```bash
go run main.go -verbose
```

Output:

```text
true
```

You can also explicitly provide:

```bash
go run main.go -verbose=true
```

or:

```bash
go run main.go -verbose=false
```

The Go documentation specifically notes that boolean flags use `-flag=false` when explicitly turning them off. :chatgpt-content-reference{index="4"}

---

# 14. `Duration`

This is extremely useful for servers and CLI applications.

```go
timeout := flag.Duration(
	"timeout",
	30*time.Second,
	"Request timeout",
)

flag.Parse()

fmt.Println(*timeout)
```

Import:

```go
import "time"
```

Run:

```bash
go run main.go -timeout=5s
```

Output:

```text
5s
```

You can use values understood by `time.ParseDuration`, such as:

```text
5s
10s
500ms
2m
1h
```

The official `flag` documentation confirms that `Duration` accepts values understood by `time.ParseDuration`. :chatgpt-content-reference{index="5"}

---

# 15. Summary of Built-in Flag Types

The main built-in constructors are:

| Function | Type | Example |
|---|---|---|
| `flag.String()` | `string` | `"hello"` |
| `flag.Bool()` | `bool` | `true` |
| `flag.Int()` | `int` | `100` |
| `flag.Int64()` | `int64` | `100000` |
| `flag.Uint()` | `uint` | `100` |
| `flag.Uint64()` | `uint64` | `100000` |
| `flag.Float64()` | `float64` | `12.5` |
| `flag.Duration()` | `time.Duration` | `5s` |

These functions return pointers containing the parsed value. :chatgpt-content-reference{index="6"}

---

# 16. The `Var` Versions

There is another important family:

```text
StringVar
BoolVar
IntVar
Int64Var
UintVar
Uint64Var
Float64Var
DurationVar
```

The difference is:

### Without `Var`

Go creates the variable for you:

```go
name := flag.String("name", "Kasra", "Name")
```

You receive:

```go
*string
```

### With `Var`

You create the variable yourself:

```go
var name string

flag.StringVar(
	&name,
	"name",
	"Kasra",
	"Name",
)
```

Then:

```go
fmt.Println(name)
```

No `*` is necessary.

The official documentation describes `StringVar` and the other `Var` functions as binding a flag directly to your variable. :chatgpt-content-reference{index="7"}

---

# 17. Comparing `String` and `StringVar`

### `String`

```go
name := flag.String("name", "Kasra", "Name")

flag.Parse()

fmt.Println(*name)
```

### `StringVar`

```go
var name string

flag.StringVar(
	&name,
	"name",
	"Kasra",
	"Name",
)

flag.Parse()

fmt.Println(name)
```

Personally, for larger CLI programs, I recommend becoming comfortable with the `Var` approach because it makes configuration structures easier to manage.

---

# 18. Multiple Flags

You can define as many flags as you want.

```go
package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String(
		"name",
		"Kasra",
		"User name",
	)

	age := flag.Int(
		"age",
		31,
		"User age",
	)

	active := flag.Bool(
		"active",
		true,
		"Whether user is active",
	)

	flag.Parse()

	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Active:", *active)
}
```

Run:

```bash
go run main.go -name=Ali -age=25 -active=false
```

Result:

```text
Name: Ali
Age: 25
Active: false
```

---

# 19. `flag.Parse()`

This is one of the most important functions.

You define:

```go
name := flag.String("name", "Kasra", "Name")
```

But Go hasn't processed the command line yet.

You need:

```go
flag.Parse()
```

Think of it as:

```text
Command line
     ↓
flag.Parse()
     ↓
Parse all flags
     ↓
Store values in variables
```

The official documentation says that after defining all flags, you call `flag.Parse()` to parse the command line. :chatgpt-content-reference{index="8"}

---

# 20. What Happens If You Don't Call `Parse()`?

Example:

```go
name := flag.String(
	"name",
	"Kasra",
	"Name",
)

fmt.Println(*name)
```

Run:

```bash
go run main.go -name=Ali
```

You will still get:

```text
Kasra
```

because parsing never happened.

Therefore:

```go
flag.Parse()
```

is normally required.

---

# 21. `flag.Args()`

Now we get to **non-flag arguments**.

Suppose:

```bash
myapp -verbose file1.txt file2.txt
```

`-verbose` is a flag.

These:

```text
file1.txt
file2.txt
```

are arguments.

You can retrieve them using:

```go
args := flag.Args()

fmt.Println(args)
```

Output:

```text
[file1.txt file2.txt]
```

The documentation defines `Args()` as returning the non-flag command-line arguments. :chatgpt-content-reference{index="9"}

---

# 22. `flag.Arg(i)`

Instead of getting all arguments:

```go
flag.Args()
```

you can get one:

```go
flag.Arg(0)
```

Example:

```bash
myapp file1.txt file2.txt
```

Then:

```go
fmt.Println(flag.Arg(0))
```

gives:

```text
file1.txt
```

And:

```go
fmt.Println(flag.Arg(1))
```

gives:

```text
file2.txt
```

If the requested argument doesn't exist, `Arg()` returns an empty string. :chatgpt-content-reference{index="10"}

---

# 23. `flag.NArg()`

Returns the number of non-flag arguments.

Example:

```bash
myapp file1.txt file2.txt file3.txt
```

```go
fmt.Println(flag.NArg())
```

Output:

```text
3
```

Very useful:

```go
if flag.NArg() == 0 {
	fmt.Println("No files provided")
}
```

---

# 24. `flag.NFlag()`

Returns the number of flags that were actually set.

Example:

```bash
myapp -verbose -name=Kasra
```

```go
fmt.Println(flag.NFlag())
```

Output:

```text
2
```

Important distinction:

```text
NFlag()
    ↓
Number of flags explicitly provided

NArg()
    ↓
Number of non-flag arguments
```

---

# 25. `flag.Parsed()`

Checks whether `flag.Parse()` has already been called.

```go
if flag.Parsed() {
	fmt.Println("Flags have been parsed")
}
```

Before:

```go
flag.Parse()
```

it is:

```text
false
```

After:

```go
flag.Parse()
```

it becomes:

```text
true
```

---

# 26. `flag.PrintDefaults()`

This is extremely useful for CLI applications.

Suppose:

```go
name := flag.String(
	"name",
	"Kasra",
	"User name",
)

age := flag.Int(
	"age",
	31,
	"User age",
)

flag.Parse()

flag.PrintDefaults()
```

You can get output similar to:

```text
  -age int
        User age (default 31)
  -name string
        User name (default "Kasra")
```

The function prints documentation for all defined command-line flags. :chatgpt-content-reference{index="11"}

---

# 27. The Built-in `-h` / `-help`

One of the nicest features is that Go automatically supports help.

Given:

```go
package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String(
		"name",
		"Kasra",
		"User name",
	)

	age := flag.Int(
		"age",
		31,
		"User age",
	)

	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*age)
}
```

Run:

```bash
go run main.go -h
```

You'll get usage information containing the available flags and their defaults.

The package has built-in handling for `-h` and `-help`. :chatgpt-content-reference{index="12"}

---

# 28. `flag.Set()`

You can programmatically set a flag.

For example:

```go
name := flag.String(
	"name",
	"Kasra",
	"User name",
)

flag.Set("name", "Ali")

fmt.Println(*name)
```

Output:

```text
Ali
```

Signature:

```go
flag.Set(name, value string) error
```

This is useful when configuration comes from somewhere other than the command line.

---

# 29. `flag.Lookup()`

This lets you find a flag by name.

Example:

```go
name := flag.String(
	"name",
	"Kasra",
	"User name",
)

flag.Parse()

f := flag.Lookup("name")

fmt.Println(f.Name)
fmt.Println(f.DefValue)
```

This is important because `*flag.Flag` contains metadata about the flag.

A `Flag` contains information such as:

```text
Name
Usage
Value
DefValue
```

The official API exposes `Lookup` specifically for retrieving a flag by name. :chatgpt-content-reference{index="13"}

---

# 30. Getting the Default Value Properly

This addresses your question directly.

Suppose:

```go
name := flag.String(
	"name",
	"Kasra",
	"User name",
)
```

There are two different concepts:

### Current value

```go
fmt.Println(*name)
```

Could be:

```text
Ali
```

if the user ran:

```bash
myapp -name=Ali
```

### Default value

You can inspect the `Flag`:

```go
f := flag.Lookup("name")

fmt.Println(f.DefValue)
```

Output:

```text
Kasra
```

So:

```go
f.DefValue
```

is the **default value as a string**.

This is different from:

```go
*f.Value
```

which represents the current value.

---

# 31. Current Value vs Default Value

This distinction is extremely important.

Suppose:

```go
name := flag.String(
	"name",
	"Kasra",
	"User name",
)

flag.Parse()

f := flag.Lookup("name")

fmt.Println("Current:", *name)
fmt.Println("Default:", f.DefValue)
```

Run:

```bash
go run main.go
```

Result:

```text
Current: Kasra
Default: Kasra
```

Run:

```bash
go run main.go -name=Ali
```

Result:

```text
Current: Ali
Default: Kasra
```

Therefore:

```text
*name
   ↓
Current value

f.DefValue
   ↓
Default value
```

---

# 32. `flag.Var()` — Custom Types

This is where `flag` becomes more powerful.

Suppose you want:

```bash
-tags=go,api,microservices
```

There isn't a built-in:

```go
flag.StringSlice()
```

in the basic API.

You can create your own type implementing `flag.Value`.

The interface requires:

```go
type Value interface {
	String() string
	Set(string) error
}
```

The official documentation demonstrates this mechanism with a custom duration-slice type. :chatgpt-content-reference{index="14"}

Example:

```go
package main

import (
	"flag"
	"fmt"
	"strings"
)

type StringSlice []string

func (s *StringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *StringSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {

	var tags StringSlice

	flag.Var(
		&tags,
		"tags",
		"Comma-separated tags",
	)

	flag.Parse()

	fmt.Println(tags)
}
```

Run:

```bash
go run main.go -tags=go,api,microservices
```

Result:

```text
[go api microservices]
```

This is an important concept when building serious CLI applications.

---

# 33. `flag.Func()`

`Func` allows you to execute your own function when a flag is provided.

Example:

```go
flag.Func(
	"name",
	"Set the name",
	func(value string) error {
		fmt.Println("Name received:", value)
		return nil
	},
)
```

Then:

```bash
myapp -name=Kasra
```

causes your function to receive:

```text
Kasra
```

The function has:

```go
func(string) error
```

as its callback.

---

# 34. `flag.BoolFunc()`

`BoolFunc` is similar but designed for boolean-style flags.

Example:

```go
flag.BoolFunc(
	"verbose",
	"Enable verbose mode",
	func(value string) error {
		fmt.Println("Verbose:", value)
		return nil
	},
)
```

Then:

```bash
myapp -verbose
```

passes:

```text
true
```

The documentation notes that `BoolFunc` was added in Go 1.21. :chatgpt-content-reference{index="15"}

---

# 35. `flag.TextVar()`

This is a more advanced option.

It allows a value implementing Go's:

```go
encoding.TextUnmarshaler
```

interface to be used as a flag.

This becomes useful when you have custom types that know how to convert themselves from text.

For example, custom configuration types.

You don't need this for normal CLI programs initially, but it's worth knowing that it exists.

---

# 36. `flag.Visit()`

`Visit()` lets you iterate over **flags that were actually set**.

Example:

```go
flag.Visit(func(f *flag.Flag) {
	fmt.Println("Flag:", f.Name)
})
```

Suppose:

```bash
myapp -name=Kasra -age=31
```

Then your callback receives those flags.

This is useful when you need to know:

> Which flags did the user explicitly provide?

This is different from `VisitAll()`.

---

# 37. `flag.VisitAll()`

`VisitAll()` iterates over **all defined flags**.

Example:

```go
flag.VisitAll(func(f *flag.Flag) {
	fmt.Println(f.Name)
})
```

Suppose you define:

```go
-name
-age
-verbose
```

Even if the user runs:

```bash
myapp -name=Kasra
```

`VisitAll()` will still visit:

```text
name
age
verbose
```

Whereas:

```go
Visit()
```

only visits:

```text
name
```

because that was explicitly provided.

---

# 38. `flag.Arg()` vs `flag.Args()`

Remember:

```go
flag.Arg(0)
```

means:

> Give me argument number 0.

While:

```go
flag.Args()
```

means:

> Give me all remaining arguments.

Example:

```bash
myapp -verbose file1.txt file2.txt
```

Then:

```go
flag.Arg(0)
```

→

```text
file1.txt
```

and:

```go
flag.Args()
```

→

```text
[file1.txt file2.txt]
```

---

# 39. Command-Line Syntax

Go accepts:

```bash
-name=Kasra
```

and:

```bash
-name Kasra
```

for non-boolean flags.

It also accepts:

```bash
--name=Kasra
```

because one or two dashes are accepted. :chatgpt-content-reference{index="16"}

So these are equivalent:

```bash
-name=Kasra
```

```bash
--name=Kasra
```

---

# 40. Boolean Syntax

For boolean:

```bash
-verbose
```

means:

```text
true
```

You can also:

```bash
-verbose=true
```

or:

```bash
-verbose=false
```

But this is not normally used:

```bash
-verbose false
```

The parser treats boolean flags differently because of command-line parsing rules. :chatgpt-content-reference{index="17"}

---

# 41. Parsing Stops at a Non-Flag Argument

This is important.

Suppose:

```bash
myapp -name=Kasra file.txt -age=31
```

The parser stops when it reaches:

```text
file.txt
```

because it is a non-flag argument.

Therefore:

```text
-name=Kasra
```

is parsed as a flag.

```text
file.txt
```

becomes an argument.

And:

```text
-age=31
```

is no longer parsed as a flag.

The official documentation explicitly describes this behavior. :chatgpt-content-reference{index="18"}

---

# 42. `--`

You can also use:

```bash
--
```

as a terminator.

Example:

```bash
myapp -name=Kasra -- -file.txt -something
```

Everything after:

```text
--
```

is treated as arguments rather than flags.

---

# 43. A Realistic CLI Example

Let's put everything together.

```go
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	name := flag.String(
		"name",
		"Guest",
		"User name",
	)

	age := flag.Int(
		"age",
		18,
		"User age",
	)

	verbose := flag.Bool(
		"verbose",
		false,
		"Enable verbose output",
	)

	timeout := flag.Duration(
		"timeout",
		30*time.Second,
		"Request timeout",
	)

	flag.Parse()

	fmt.Println("========== Configuration ==========")

	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Verbose:", *verbose)
	fmt.Println("Timeout:", *timeout)

	fmt.Println()

	fmt.Println("========== Arguments ==========")

	fmt.Println("Argument count:", flag.NArg())
	fmt.Println("Arguments:", flag.Args())

	fmt.Println()

	fmt.Println("========== Flags ==========")

	fmt.Println("Flag count:", flag.NFlag())
}
```

Run:

```bash
go run main.go
```

You get:

```text
========== Configuration ==========
Name: Guest
Age: 18
Verbose: false
Timeout: 30s

========== Arguments ==========
Argument count: 0
Arguments: []

========== Flags ==========
Flag count: 0
```

Now:

```bash
go run main.go -name=Kasra -age=31 -verbose -timeout=5s file.txt
```

You get approximately:

```text
========== Configuration ==========
Name: Kasra
Age: 31
Verbose: true
Timeout: 5s

========== Arguments ==========
Argument count: 1
Arguments: [file.txt]

========== Flags ==========
Flag count: 4
```

---

# 44. `FlagSet` — The Next Important Concept

So far we've used:

```go
flag.String()
flag.Int()
flag.Parse()
```

These operate on the package's default `FlagSet`.

But Go also allows you to create your own:

```go
flag.NewFlagSet()
```

This becomes very important for **CLI subcommands**.

For example:

```bash
myapp backup
myapp restore
myapp config
```

Each command can have its own flags.

The official documentation describes `FlagSet` as a way to create independent sets of flags, particularly useful for CLI subcommands. :chatgpt-content-reference{index="19"}

---

# 45. Creating a `FlagSet`

```go
fs := flag.NewFlagSet(
	"backup",
	flag.ExitOnError,
)
```

Then:

```go
source := fs.String(
	"source",
	"data",
	"Source directory",
)
```

And:

```go
fs.Parse(os.Args[2:])
```

Notice:

```go
flag.Parse()
```

became:

```go
fs.Parse(...)
```

---

# 46. Why `FlagSet` Is Useful

Imagine:

```text
myapp backup
myapp restore
myapp delete
```

You could have:

```text
backup:
    -source
    -destination

restore:
    -file

delete:
    -force
```

This is much cleaner than putting every possible option into one giant flag list.

---

# 47. `NewFlagSet`

Signature:

```go
flag.NewFlagSet(
	name,
	errorHandling,
)
```

Example:

```go
fs := flag.NewFlagSet(
	"backup",
	flag.ExitOnError,
)
```

There are three important error-handling modes:

```go
flag.ContinueOnError
```

```go
flag.ExitOnError
```

```go
flag.PanicOnError
```

---

# 48. `ContinueOnError`

```go
fs := flag.NewFlagSet(
	"backup",
	flag.ContinueOnError,
)
```

Parsing returns an error instead of automatically terminating the program.

Useful when you want to handle errors yourself.

---

# 49. `ExitOnError`

```go
fs := flag.NewFlagSet(
	"backup",
	flag.ExitOnError,
)
```

If parsing fails, the program exits.

This is commonly convenient for simple CLI applications.

---

# 50. `PanicOnError`

```go
fs := flag.NewFlagSet(
	"backup",
	flag.PanicOnError,
)
```

Parsing errors cause a panic.

Less common for normal CLI applications.

---

# 51. `FlagSet.Parse`

With a custom `FlagSet`:

```go
err := fs.Parse(os.Args[1:])
```

Unlike package-level:

```go
flag.Parse()
```

`FlagSet.Parse()` returns an error:

```go
error
```

This gives you more control.

---

# 52. `FlagSet` Methods

Almost everything we've learned has an equivalent on `FlagSet`.

For example:

```go
fs.String()
fs.Int()
fs.Bool()
fs.Float64()
fs.Duration()
```

and:

```go
fs.StringVar()
fs.IntVar()
fs.BoolVar()
```

and:

```go
fs.Parse()
fs.Args()
fs.Arg()
fs.NArg()
fs.NFlag()
fs.Lookup()
fs.Set()
fs.Visit()
fs.VisitAll()
fs.PrintDefaults()
```

The official documentation lists these as methods of `FlagSet`. :chatgpt-content-reference{index="20"}

---

# 53. Important `FlagSet` Methods

| Method | Purpose |
|---|---|
| `NewFlagSet()` | Create independent flags |
| `String()` | Define string flag |
| `StringVar()` | Bind string flag to variable |
| `Int()` | Define integer flag |
| `IntVar()` | Bind integer flag |
| `Bool()` | Define boolean flag |
| `BoolVar()` | Bind boolean flag |
| `Float64()` | Define float flag |
| `Duration()` | Define duration flag |
| `Var()` | Custom flag type |
| `Parse()` | Parse arguments |
| `Args()` | Get non-flag arguments |
| `Arg()` | Get one argument |
| `NArg()` | Number of arguments |
| `NFlag()` | Number of supplied flags |
| `Lookup()` | Find a flag |
| `Set()` | Set a flag programmatically |
| `Visit()` | Iterate over supplied flags |
| `VisitAll()` | Iterate over all flags |
| `PrintDefaults()` | Print help/defaults |
| `SetOutput()` | Change help/error output |
| `Output()` | Get current output |
| `Name()` | Get FlagSet name |
| `Parsed()` | Check whether parsed |
| `ErrorHandling()` | Get error-handling mode |
| `Init()` | Reinitialize a FlagSet |

---

# 54. The Most Important Functions to Learn First

Don't try to memorize every function immediately.

For your current CLI learning path, learn these in this order:

### Level 1 — Essential

```go
flag.String()
flag.Int()
flag.Bool()
flag.Parse()
```

Then:

```go
flag.Args()
flag.Arg()
flag.NArg()
```

### Level 2 — Important

```go
flag.Duration()
flag.Float64()
flag.Int64()
flag.StringVar()
flag.IntVar()
flag.BoolVar()
```

Then:

```go
flag.PrintDefaults()
flag.Lookup()
flag.NFlag()
flag.Parsed()
```

### Level 3 — Advanced

```go
flag.Var()
flag.Func()
flag.BoolFunc()
flag.Visit()
flag.VisitAll()
```

Then:

```go
flag.FlagSet
flag.NewFlagSet()
```

### Level 4 — Specialized

```go
flag.TextVar()
flag.SetOutput()
flag.Output()
flag.Init()
flag.UnquoteUsage()
```

---

# 55. One Very Important Mental Model

Think about the whole package like this:

```text
                COMMAND LINE
                     │
                     ▼
              ┌──────────────┐
              │ flag package │
              └──────┬───────┘
                     │
              flag.Parse()
                     │
          ┌──────────┴──────────┐
          │                     │
          ▼                     ▼
       FLAGS                 ARGUMENTS
          │                     │
          │                     │
   -name=Kasra             file.txt
   -age=31                 image.jpg
   -verbose
          │                     │
          ▼                     ▼
   flag variables          flag.Args()
```

And there are two types of information:

```text
FLAG
-name=Kasra
   │
   ├── current value → Kasra
   └── default value → Guest
```

That's the key idea behind the entire package.

---

# 56. Complete Example Showing Defaults

This example specifically demonstrates the difference between **default values** and **current values**:

```go
package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String(
		"name",
		"Guest",
		"User name",
	)

	age := flag.Int(
		"age",
		18,
		"User age",
	)

	verbose := flag.Bool(
		"verbose",
		false,
		"Enable verbose mode",
	)

	flag.Parse()

	fmt.Println("===== Current Values =====")

	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Verbose:", *verbose)

	fmt.Println()

	fmt.Println("===== Default Values =====")

	nameFlag := flag.Lookup("name")
	ageFlag := flag.Lookup("age")
	verboseFlag := flag.Lookup("verbose")

	fmt.Println("Name:", nameFlag.DefValue)
	fmt.Println("Age:", ageFlag.DefValue)
	fmt.Println("Verbose:", verboseFlag.DefValue)
}
```

Run:

```bash
go run main.go
```

Output:

```text
===== Current Values =====
Name: Guest
Age: 18
Verbose: false

===== Default Values =====
Name: Guest
Age: 18
Verbose: false
```

Now:

```bash
go run main.go -name=Kasra -age=31 -verbose
```

Output:

```text
===== Current Values =====
Name: Kasra
Age: 31
Verbose: true

===== Default Values =====
Name: Guest
Age: 18
Verbose: false
```

This is probably the **most important example for your question about default values**.

---

# 57. One More Important Detail: "Was the Default Used?"

Sometimes you don't actually want to know the default value.

You want to know:

> Did the user explicitly provide this flag?

For that, `Visit()` is useful.

Example:

```go
flag.Visit(func(f *flag.Flag) {
	fmt.Println("User provided:", f.Name)
})
```

If you run:

```bash
myapp -name=Kasra
```

you'll see:

```text
User provided: name
```

But `age` won't appear.

That's because:

```text
Lookup()
    → find information about a flag

DefValue
    → default value

Visit()
    → flags explicitly supplied by user

VisitAll()
    → all defined flags
```

That distinction becomes very useful when you build real CLI configuration systems.

---

# 58. What I Recommend You Learn From `flag`

Since you're learning Go specifically toward **CLI → APIs → backend → microservices**, I would learn `flag` in this sequence:

```text
1. String
2. Int
3. Bool
4. Parse
5. StringVar / IntVar / BoolVar
6. Args / Arg / NArg
7. Duration
8. Float64 / Int64 / Uint
9. PrintDefaults
10. Lookup
11. NFlag
12. Visit / VisitAll
13. Var
14. Func
15. FlagSet
16. NewFlagSet
17. CLI subcommands
```

You don't need to memorize all of `flag` before moving forward.

The **core mental model** you should have now is:

```text
Define
   ↓
flag.String(...)
flag.Int(...)
flag.Bool(...)
   ↓
Parse
   ↓
flag.Parse()
   ↓
Use
   ↓
*name
*age
*verbose
   ↓
Arguments
   ↓
flag.Args()
```

And for defaults:

```text
flag.String("name", "Guest", ...)
                    ↑
                    │
               default value
```

Current value:

```go
*name
```

Default value:

```go
flag.Lookup("name").DefValue
```

Explicitly supplied flags:

```go
flag.Visit(...)
```

All defined flags:

```go
flag.VisitAll(...)
```

That covers the practical API surface of the current standard `flag` package, including the package-level functions and the corresponding `FlagSet` methods documented by Go. :chatgpt-content-reference{index="21"}

[Official Go `flag` documentation](https://pkg.go.dev/flag?utm_source=chatgpt.com)
*/
