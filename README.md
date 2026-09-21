# Go CLI / CMD / Command Prompt Projects

> **First Version of GoLang CLI / CMD / Command Prompt Projects** 🚀

A curated collection of command-line applications, tools, and experiments built with **Go (Golang)**. This repository is a learning journey and a playground for everything related to building powerful, fast, and cross-platform CLI tools.

---

## 📖 Table of Contents

- [Introduction](#-introduction)
- [What is Go?](#-what-is-go)
- [What is a CLI?](#-what-is-a-cli)
- [What is CMD?](#-what-is-cmd)
- [What is a Command Prompt?](#-what-is-a-command-prompt)
- [Project Structure](#-project-structure)
- [Related Packages](#-related-packages)
- [Getting Started](#-getting-started)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🌟 Introduction

This repository is the **first version** of a long-term project focused on building CLI (Command-Line Interface) applications using **Go**. It contains a growing set of tools, examples, and experiments that demonstrate how to build fast, portable, and user-friendly command-line programs.

Whether you are:
- A **beginner** learning Go and CLI development,
- An **intermediate** developer looking for reference implementations,
- Or a **contributor** wanting to add your own tool,

...this repo is for you.

The goal is to cover the entire lifecycle of a CLI project: from reading arguments and flags, to formatting output, handling user input, adding colors, and finally distributing the binary.

---

## 🐹 What is Go?

**Go** (also called **Golang**) is an open-source, statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It was released in 2009.

Key characteristics:
- **Fast compilation** and **fast execution** (compiled to a single native binary).
- **Simple syntax** — easy to learn, hard to misuse.
- **Built-in concurrency** with goroutines and channels.
- **Cross-platform** — compile for Linux, macOS, Windows, and more from one machine.
- **Rich standard library** — especially strong for networking, file I/O, and CLI tools.
- **Garbage collected** but with low latency.

Because Go compiles to a **single static binary** with no runtime dependencies, it is one of the best languages for building CLI tools that users can just download and run.

Official site: [https://go.dev](https://go.dev)

---

## 💻 What is a CLI?

**CLI** stands for **Command-Line Interface**. It is a way for users to interact with a program by typing text commands into a terminal or shell, instead of clicking buttons in a graphical interface.

Examples of popular CLIs:
- `git`
- `docker`
- `kubectl`
- `npm`
- `go`

A CLI program typically:
1. Reads **arguments** and **flags** (e.g., `myapp --name Ali -v`).
2. Reads **input** from the user (via `stdin`) or from files.
3. Processes the input.
4. Writes **output** to the terminal (`stdout`) or to a file.
5. Returns an **exit code** (`0` for success, non-zero for errors).

CLIs are preferred by developers because they are **fast**, **scriptable**, **automatable**, and **composable** (you can pipe the output of one into another).

---

## 🪟 What is CMD?

**CMD** can refer to two related things:

1. **`cmd` (as a Go package/folder convention):**
   In Go projects, it is common to place the `main` packages of executables inside a `cmd/` directory. Each subfolder represents a separate binary. For example:
   ```
   myproject/
   ├── cmd/
   │   ├── mytool/
   │   │   └── main.go
   │   └── anothertool/
   │       └── main.go
   ├── pkg/
   └── go.mod
   ```
   This structure lets one repository produce **multiple CLI binaries**.

2. **CMD as "Command":**
   A "command" is any executable instruction you type into a shell. In Go, a command is usually a program with a `main()` function.

---

## 🖥️ What is a Command Prompt?

A **Command Prompt** (or **terminal**, **shell**, **console**) is the program that lets you type commands and see their output.

Examples across operating systems:
- **Windows:** `cmd.exe`, **PowerShell**, **Windows Terminal**
- **macOS / Linux:** `bash`, `zsh`, `fish`, **Terminal**, **iTerm2**
- **Cross-platform:** **Windows Terminal**, **Alacritty**, **Kitty**

When you run a Go CLI program, it executes inside one of these shells. The shell is responsible for passing arguments, handling pipes, and rendering output — including colors and styles.

---

## 📂 Project Structure

A typical layout used in this repository:

```
.
├── cmd/                # Entry points for each CLI binary
│   └── <tool-name>/
│       └── main.go
├── internal/           # Private packages (not importable by other modules)
├── pkg/                # Reusable public packages
├── go.mod
├── go.sum
├── Makefile            # Optional build helpers
└── README.md
```

This follows the official [Go project layout](https://github.com/golang-standards/project-layout) conventions.

---

## 📦 Related Packages

Here are the most useful Go packages for building CLI applications, organized by category.

### 1. Standard Library (No installation needed)
| Package | Purpose |
|---|---|
| `os` | Access arguments (`os.Args`), environment variables, files |
| `flag` | Parse command-line flags (`-name`, `-v`) |
| `fmt` | Formatted printing and scanning |
| `bufio` | Buffered reading (line-by-line input, large files) |
| `io` | Core interfaces (`Reader`, `Writer`) for testable I/O |
| `path/filepath` | Cross-platform file path manipulation |
| `os/exec` | Run external commands |
| `encoding/json` | Encode/decode JSON |
| `context` | Cancellation, timeouts, request-scoped values |

### 2. CLI Frameworks
| Package | Description |
|---|---|
| [`spf13/cobra`](https://github.com/spf13/cobra) | The most popular CLI framework. Used by `kubectl`, `docker`, `hugo`. |
| [`urfave/cli`](https://github.com/urfave/cli) | Simple and idiomatic CLI framework. |
| [`alecthomas/kong`](https://github.com/alecthomas/kong) | Struct-based CLI parser. |

### 3. Terminal UI
| Package | Description |
|---|---|
| [`charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea) | Elm-architecture TUI framework. |
| [`charmbracelet/lipgloss`](https://github.com/charmbracelet/lipgloss) | Style and layout for terminal output. |
| [`charmbracelet/bubbles`](https://github.com/charmbracelet/bubbles) | Reusable TUI components (spinners, lists, inputs). |

### 4. Configuration
| Package | Description |
|---|---|
| [`spf13/viper`](https://github.com/spf13/viper) | Config files (YAML, TOML, JSON, ENV) + flags. |
| [`joho/godotenv`](https://github.com/joho/godotenv) | Load `.env` files. |

### 5. Output Formatting
| Package | Description |
|---|---|
| [`olekukonko/tablewriter`](https://github.com/olekukonko/tablewriter) | Pretty ASCII tables. |
| [`fatih/color`](https://github.com/fatih/color) | ANSI colors and styles. |
| [`logrusorgru/aurora`](https://github.com/logrusorgru/aurora) | Alternative color library. |
| `encoding/json`, `gopkg.in/yaml.v3` | Structured output. |

### 6. Distribution
| Tool | Description |
|---|---|
| `go build` | Build a single binary. |
| [`goreleaser/goreleaser`](https://github.com/goreleaser/goreleaser) | Automate cross-platform releases. |
| **Docker** | Ship your CLI in a container. |

---

## 🚀 Getting Started

### Prerequisites
- **Go 1.21+** — [Download here](https://go.dev/dl/)
- A terminal (bash, zsh, PowerShell, etc.)
- (Optional) `make` for build shortcuts

### Clone & Run
```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo

# Build all binaries
go build ./...

# Run a specific tool
go run ./cmd/<tool-name> --help
```

### Build a Single Binary
```bash
go build -o bin/mytool ./cmd/mytool
./bin/mytool --help
```

### Cross-Compile (Example: Linux → Windows)
```bash
GOOS=windows GOARCH=amd64 go build -o bin/mytool.exe ./cmd/mytool
```

---

## 🤝 Contributing

Just Send Ideas For this Or fix Bugs.

---

## ⭐ Show Your Support

If you find this repository useful, please give it a **star** ⭐ — it helps others discover it too!

Happy hacking! 🐹✨
