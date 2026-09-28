package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Program Is Started")

	userID := 10

	log.Print("Loading User: ", userID)
	// Print joins values.
	// Note: log.Print adds a space between arguments when neither
	// argument is a string that already ends with a space.

	log.Printf("Loading User: %d", userID)
	// Printf uses a format string:
	// %d  -> integer
	// %s  -> string
	// %v  -> general value
	// %+v -> value with more detail, especially structs

	fmt.Println("\n--- Log Flags ---")

	fmt.Println("log.Ldate:         ", log.Ldate)         // 1
	fmt.Println("log.Ltime:         ", log.Ltime)         // 2
	fmt.Println("log.Lmicroseconds: ", log.Lmicroseconds) // 4
	fmt.Println("log.Llongfile:     ", log.Llongfile)     // 8
	fmt.Println("log.Lshortfile:    ", log.Lshortfile)    // 16
	fmt.Println("log.LUTC:          ", log.LUTC)          // 32
	fmt.Println("log.Lmsgprefix:    ", log.Lmsgprefix)    // 64
	fmt.Println("log.LstdFlags:     ", log.LstdFlags)     // 3

	fmt.Println("\n--- What The Flags Do ---")

	fmt.Println("Ldate         -> Date: YYYY/MM/DD")
	fmt.Println("Ltime         -> Time: HH:MM:SS")
	fmt.Println("Lmicroseconds -> Time + microseconds")
	fmt.Println("Llongfile     -> Full file path + line number")
	fmt.Println("Lshortfile    -> File name + line number")
	fmt.Println("LUTC          -> Use UTC time")
	fmt.Println("Lmsgprefix    -> Put prefix before message")
	fmt.Println("LstdFlags     -> Ldate | Ltime")

	fmt.Println("\n--- Examples ---")

	// Default logger.
	log.SetFlags(log.LstdFlags)
	log.Println("Using LstdFlags")

	// Date only.
	log.SetFlags(log.Ldate)
	log.Println("Using Ldate")

	// Time only.
	log.SetFlags(log.Ltime)
	log.Println("Using Ltime")

	// Date + time.
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Using Ldate | Ltime")

	// Date + time + microseconds.
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("Using Ldate | Ltime | Lmicroseconds")

	// Date + time + source file and line number.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Using LstdFlags | Lshortfile")

	// Full source file path + line number.
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("Using LstdFlags | Llongfile")

	// UTC time.
	log.SetFlags(log.LstdFlags | log.LUTC)
	log.Println("Using LstdFlags | LUTC")

	// ------------------------------------------------
	// Creating a custom Logger
	// ------------------------------------------------

	fmt.Println("\n--- Custom Logger ---")

	logger := log.New(
		os.Stdout,
		"APP: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	logger.Println("Custom logger message")

	// Method 2, create with set them separately
	log.SetPrefix("API: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Server is listening on port 8080")

	myLogger := log.New(os.Stdout, "myApp Says: ", log.Ldate|log.Ltime|log.Lshortfile)
	myLogger.Println("Database Connection Problem Code: 404!!!:)))")
}

/*
| Flag              | Adds to a log line                                       |
| ----------------- | -------------------------------------------------------- |
| log.Ldate         | Local date, such as 2026/09/28                           |
| log.Ltime         | Local time, such as 09:28:00                             |
| log.Lmicroseconds | Microsecond precision; includes time                     |
| log.Llongfile     | Full source path and line number                         |
| log.Lshortfile    | Only file name and line number                           |
| log.LUTC          | Uses UTC for date/time                                   |
| log.Lmsgprefix    | Places your custom prefix immediately before the message |
| log.LstdFlags     | Default combination: Ldate \| Ltime                      |
*/

/*
log.SetFlags(0)
log.Println("Connected")
* With flag value 0, the logger adds no date, time, or file information.
*/

/*
infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("Application started")
	errorLog.Println("Database connection failed")
	Common writers include:

os.Stdout for regular console output.
os.Stderr for errors and diagnostics.
An *os.File created with os.OpenFile for file logging.
bytes.Buffer in tests, so you can inspect what was logged.
A Logger is safe to use from several goroutines: it serializes writes so one log entry is not mixed halfway through another entry.
*/

/*
* Fatal vs Panic
These functions log a message, but they also stop normal execution:

log.Fatal("Cannot start without configuration")
Fatal, Fatalf, and Fatalln write the message and then call os.Exit(1). Deferred functions do not run after os.Exit, so avoid Fatal deep inside reusable functions where cleanup may matter. It is best reserved for unrecoverable startup failures in main.

log.Panic("Unexpected invalid state")
Panic, Panicf, and Panicln write the message and then call panic. Unlike Fatal, a panic can be recovered with recover, and deferred functions run as the stack unwinds.

For normal, recoverable errors, return an error instead of logging it repeatedly:

func readConfig() error {
	// Return the error to the caller.
	return nil
}
A useful rule: the layer that can make the decision should usually log the error—often main or an HTTP handler—not every function in the call chain.
*/

/*
## Daily log function

Good goal. Use one function that:

1. Creates `logs/` if it does not exist.
2. Generates a filename from today’s date, such as `log-2026-09-28.txt`.
3. Opens that file in **append** mode.
4. Returns a reusable `*log.Logger`.

`os.MkdirAll` safely does nothing if the directory already exists, while `os.OpenFile` with `O_CREATE` creates a missing file and `O_APPEND` writes new data at its end.

```go
package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func NewDailyLogger() (*log.Logger, *os.File, error) {
	const logsDir = "logs"

	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		return nil, nil, err
	}

	today := time.Now().Format("2006-01-02")

	fileName := "log-" + today + ".txt"

	filePath := filepath.Join(logsDir, fileName)

	file, err := os.OpenFile(
		filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(
		file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	return logger, file, nil
}
```

The special time layout `2006-01-02` means year-month-day in Go, so it produces filenames that sort naturally by date. The `log.New` function creates a logger that writes to the file, and `Ldate`, `Ltime`, and `Lshortfile` add useful context to every entry. [pkg.go](https://pkg.go.dev/log)

## Using it in `main`

Put the code above in a file such as `logger/logger.go`. Then use it like this:

```go
package main

import (
	"log"

	"your-project/logger"
)

func main() {
	appLog, logFile, err := logger.NewDailyLogger()
	if err != nil {
		log.Fatal("Could not create log file:", err)
	}

	defer logFile.Close()

	appLog.Println("Application started")

	userID := 42

	appLog.Printf("User %d logged in", userID)

	appLog.Println("Application finished")
}
```

This creates a structure like:

```text
your-project/
├── main.go
├── logger/
│   └── logger.go
└── logs/
    └── log-2026-09-28.txt
```

Every later call on the same day appends to that day’s file; when the calendar date changes and the program starts again, the function selects a different filename automatically. `defer logFile.Close()` is important because it closes the operating-system file handle when `main` ends.

## Example file content

```text
INFO: 2026/09/28 10:05:13 main.go:18: Application started
INFO: 2026/09/28 10:05:13 main.go:22: User 42 logged in
INFO: 2026/09/28 10:05:13 main.go:24: Application finished
```

One important detail: this version chooses the file **when `NewDailyLogger()` runs**. For most CLI programs and web servers restarted regularly, that is enough. A server that stays running across midnight needs a slightly more advanced logger that checks the date before every write and switches files automatically.
*/
