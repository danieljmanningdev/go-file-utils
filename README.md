# gofileutils
A lightweight, developer-friendly Go utility package designed to simplify common file operations, text scanning, and JSON parsing. It cuts down on boilerplate code by wrapping standard library I/O routines into clean, reusable functions.

## Installation
Add the package to your Go project via GitHub:
```
go get github.com/danieljmanningdev/gofileutils
```

## Features & Usage
The package introduces a custom filepath type for paths, alongside helper functions that handle file opening, closing, scanning, and JSON marshaling automatically.

## 1. Opening Files
```go
package main

import (
    "github.com/danieljmanningdev/gofileutils"
)

func main() {
    path := gofileutils.Filepath("example.txt")
    file := gofileutils.OpenFile(path)
    defer file.Close()
}
```

## 2. Reading Text Files Line-by-Line
```go
package main

import (
    "fmt"
    "github.com/danieljmanningdev/gofileutils"
)

func main() {
    path := gofileutils.Filepath("notes.txt")
    lines := gofileutils.ReadLines(path)

    for i, line := range lines {
        fmt.Printf("Line %d: %s\n", i+1, line)
    }
}
```

## 3. Writing Lines to a File
```go
package main

import (
    "github.com/danieljmanningdev/gofileutils"
)

func main() {
    path := gofileutils.Filepath("output.txt")
    lines := []string{"First line", "Second line", "Third line"}
    
    gofileutils.WriteLines(path, lines)
}
```

## 4. Reading and Writing JSON
```go
package main

import (
    "fmt"
    "github.com/danieljmanningdev/gofileutils"
)

type Config struct {
    App    string `json:"app"`
    Active bool   `json:"active"`
}

func main() {
    path := gofileutils.Filepath("config.json")

    // Writing JSON
    cfg := Config{App: "FileUtility", Active: true}
    gofileutils.WriteJSON(path, cfg)

    // Reading JSON
    var loadedConfig Config
    gofileutils.ReadJSON(path, &loadedConfig)
    
    fmt.Printf("Loaded Config: %+v\n", loadedConfig)
}
```

