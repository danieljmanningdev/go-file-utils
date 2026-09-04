# go-file-utils

[![CI](https://github.com/danieljmanningdev/go-file-utils/actions/workflows/ci.yml/badge.svg)](https://github.com/danieljmanningdev/go-file-utils/actions/workflows/ci.yml)

`go-file-utils` is a lightweight Go package providing small, reusable helpers for common file and JSON operations.

It wraps repetitive standard-library file handling while keeping error handling in the caller's control.

## Installation

Add the package to your project:

```bash
go get github.com/danieljmanningdev/go-file-utils
```

Then import it:

```go
import "github.com/danieljmanningdev/go-file-utils"
```

## Features

* Check whether a file exists
* Open files using standard Go error handling
* Read text files line by line
* Write slices of strings as individual lines
* Decode JSON directly into structs, maps, or other Go values
* Write pretty-printed JSON to files
* Uses ordinary `string` paths with no custom path types
* Returns errors to the caller rather than forcing panic-based handling

## Usage

### Check Whether a File Exists

`FileExists` returns `true` when the path exists and points to a file.

```go
package main

import (
	"fmt"

	gofileutils "github.com/danieljmanningdev/go-file-utils"
)

func main() {
	if gofileutils.FileExists("config.json") {
		fmt.Println("File exists")
		return
	}

	fmt.Println("File not found")
}
```

### Open a File

`OpenFile` opens a file for reading and returns the file together with any error encountered.

```go
package main

import (
	"log"

	gofileutils "github.com/danieljmanningdev/go-file-utils"
)

func main() {
	file, err := gofileutils.OpenFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
```

### Read a File Line by Line

`ReadLines` reads a text file using a scanner and returns each line as an element in a slice.

```go
package main

import (
	"fmt"
	"log"

	gofileutils "github.com/danieljmanningdev/go-file-utils"
)

func main() {
	lines, err := gofileutils.ReadLines("notes.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		fmt.Printf("Line %d: %s\n", i+1, line)
	}
}
```

### Write Lines to a File

`WriteLines` creates or overwrites a file and writes each string as a separate line.

```go
package main

import (
	"log"

	gofileutils "github.com/danieljmanningdev/go-file-utils"
)

func main() {
	lines := []string{
		"First line",
		"Second line",
		"Third line",
	}

	if err := gofileutils.WriteLines("output.txt", lines); err != nil {
		log.Fatal(err)
	}
}
```

### Read JSON

`ReadJSON` opens a JSON file and decodes its contents directly into a target Go value.

The target should normally be passed as a pointer.

```go
package main

import (
	"fmt"
	"log"

	gofileutils "github.com/danieljmanningdev/go-file-utils"
)

type Config struct {
	App    string `json:"app"`
	Active bool   `json:"active"`
}

func main() {
	var config Config

	if err := gofileutils.ReadJSON("config.json", &config); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded config: %+v\n", config)
}
```

`ReadJSON` can also decode into maps or other compatible Go values:

```go
var data map[string]any

if err := gofileutils.ReadJSON("data.json", &data); err != nil {
	log.Fatal(err)
}
```

### Write JSON

`WriteJSON` marshals a Go value as indented JSON and writes it to the specified file.

```go
package main

import (
	"log"

	gofileutils "github.com/danieljmanningdev/go-file-utils"
)

type Config struct {
	App    string `json:"app"`
	Active bool   `json:"active"`
}

func main() {
	config := Config{
		App:    "FileUtility",
		Active: true,
	}

	if err := gofileutils.WriteJSON("config.json", config); err != nil {
		log.Fatal(err)
	}
}
```

## API

```go
func FileExists(path string) bool

func OpenFile(path string) (*os.File, error)

func ReadLines(path string) ([]string, error)

func WriteLines(path string, lines []string) error

func ReadJSON(path string, target any) error

func WriteJSON(path string, value any) error
```

## Error Handling

Operations that can fail return errors to the caller.

This allows applications to decide whether an error should be logged, returned, retried, displayed to a user, or treated as fatal.

```go
lines, err := gofileutils.ReadLines("notes.txt")
if err != nil {
	log.Fatal(err)
}
```

The package does not require callers to use panic-based error handling.

## Project Structure

The package is intentionally small and grouped by responsibility:

```text
.
├── file.go
├── json.go
├── gofileutils_test.go
├── go.mod
├── LICENSE
└── README.md
```

`file.go` contains general file and line-based helpers.

`json.go` contains JSON encoding and decoding helpers.

## Design

`go-file-utils` is intended to provide a small convenience layer over common Go standard-library file operations.

The package deliberately avoids introducing unnecessary abstractions around file paths or error handling. Functions accept ordinary string paths and return errors where appropriate.

It is suitable for use in command-line tools, web applications, scripts, utilities, and other Go programs that repeatedly perform basic file or JSON operations.

## License

This project is open source under the terms of the [MIT License](LICENSE).
