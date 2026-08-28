# go-file-utils

[![CI](https://github.com/danieljmanningdev/go-file-utils/actions/workflows/ci.yml/badge.svg)](https://github.com/danieljmanningdev/go-file-utils/actions/workflows/ci.yml)

`go-file-utils` is a lightweight Go utility package for common file operations, line-based text handling, and JSON encoding and decoding.

It wraps repetitive standard-library file I/O into small, reusable helpers while keeping error handling in the caller's control.

## Installation

Add the package to your Go project:

```bash
go get github.com/danieljmanningdev/go-file-utils
```

Then import it where needed:

```go
import "github.com/danieljmanningdev/go-file-utils"
```

## Features

- Custom `Filepath` type for file paths
- Open files with standard Go error handling
- Read text files line by line
- Write slices of strings as lines
- Read JSON directly into structs, maps, or other values
- Write pretty-printed JSON to files
- Check whether a file exists
- Optional panic-on-error helper pattern for scripts or applications

## Error Handling

Most operations return errors rather than panicking internally. This allows the calling application to decide how failures should be handled.

For cases where an error should be treated as unrecoverable, the package also provides Must.

For example:

```go
lines, err := gofileutils.ReadLines("notes.txt")
if err != nil {
    log.Fatal(err)
}
```

If your application deliberately treats an error as unrecoverable, you can optionally use a helper such as:

```go
func Must(err error) {
    if err != nil {
        panic(err)
    }
}
```

This keeps panic behaviour optional rather than forcing it on every user of the package.

## Usage

### Opening a File

`OpenFile` returns both the opened file and any error encountered.

```go
package main

import (
    "log"

    "github.com/danieljmanningdev/go-file-utils"
)

func main() {
    path := gofileutils.Filepath("example.txt")

    file, err := gofileutils.OpenFile(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
}
```

### Reading a Text File Line by Line

`ReadLines` reads a text file using a scanner and returns its contents as a slice of strings.

```go
package main

import (
    "fmt"
    "log"

    "github.com/danieljmanningdev/go-file-utils"
)

func main() {
    path := gofileutils.Filepath("notes.txt")

    lines, err := gofileutils.ReadLines(path)
    if err != nil {
        log.Fatal(err)
    }

    for i, line := range lines {
        fmt.Printf("Line %d: %s\n", i+1, line)
    }
}
```

### Writing Lines to a File

`WriteLines` creates or overwrites a file and writes each string as a separate line.

```go
package main

import (
    "log"

    "github.com/danieljmanningdev/go-file-utils"
)

func main() {
    path := gofileutils.Filepath("output.txt")

    lines := []string{
        "First line",
        "Second line",
        "Third line",
    }

    if err := gofileutils.WriteLines(path, lines); err != nil {
        log.Fatal(err)
    }
}
```

### Reading JSON

`ReadJSON` decodes JSON from a file directly into a target value.

```go
package main

import (
    "fmt"
    "log"

    "github.com/danieljmanningdev/go-file-utils"
)

type Config struct {
    App    string `json:"app"`
    Active bool   `json:"active"`
}

func main() {
    path := gofileutils.Filepath("config.json")

    var config Config

    if err := gofileutils.ReadJSON(path, &config); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Loaded config: %+v\n", config)
}
```

### Writing JSON

`WriteJSON` marshals a Go value as indented JSON and writes it to a file.

```go
package main

import (
    "log"

    "github.com/danieljmanningdev/go-file-utils"
)

type Config struct {
    App    string `json:"app"`
    Active bool   `json:"active"`
}

func main() {
    path := gofileutils.Filepath("config.json")

    config := Config{
        App:    "FileUtility",
        Active: true,
    }

    if err := gofileutils.WriteJSON(path, config); err != nil {
        log.Fatal(err)
    }
}
```

### Checking Whether a File Exists

`FileExists` returns `true` when the path exists and points to a file.

```go
package main

import (
    "fmt"

    "github.com/danieljmanningdev/go-file-utils"
)

func main() {
    path := gofileutils.Filepath("config.json")

    if gofileutils.FileExists(path) {
        fmt.Println("File is present!")
        return
    }

    fmt.Println("File not found.")
}
```

## API Overview

```go
type Filepath string

func OpenFile(f Filepath) (*os.File, error)
func ReadLines(f Filepath) ([]string, error)
func WriteLines(f Filepath, lines []string) error
func ReadJSON(f Filepath, v any) error
func WriteJSON(f Filepath, v any) error
func FileExists(f Filepath) bool
```

## Design

The package intentionally keeps its API small and focused on file-related utilities.

Functions return errors where callers may reasonably need to decide how to respond. This makes the package suitable for use in command-line tools, web applications, scripts, and other Go programs without forcing panic-based error handling.

## License

This project is open source under the terms of the [MIT License](LICENSE).
