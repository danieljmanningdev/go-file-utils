package gofileutils

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFileExists(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "test.txt"))

	if err := os.WriteFile(string(filePath), []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	if !FileExists(filePath) {
		t.Fatal("FileExists() = false, want true")
	}

	missingPath := Filepath(filepath.Join(dir, "missing.txt"))

	if FileExists(missingPath) {
		t.Fatal("FileExists() = true for missing file, want false")
	}

	dirPath := Filepath(dir)

	if FileExists(dirPath) {
		t.Fatal("FileExists() = true for directory, want false")
	}
}

func TestOpenFile(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "test.txt"))

	if err := os.WriteFile(string(filePath), []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	file, err := OpenFile(filePath)
	if err != nil {
		t.Fatalf("OpenFile() error = %v", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		t.Fatalf("file.Stat() error = %v", err)
	}

	if info.Name() != "test.txt" {
		t.Fatalf("opened file name = %q, want %q", info.Name(), "test.txt")
	}
}

func TestOpenFileMissingFile(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "missing.txt"))

	file, err := OpenFile(filePath)

	if err == nil {
		if file != nil {
			file.Close()
		}

		t.Fatal("OpenFile() error = nil, want error")
	}

	if file != nil {
		t.Fatal("OpenFile() returned non-nil file for missing path")
	}
}

func TestWriteLinesAndReadLines(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "lines.txt"))

	want := []string{
		"First line",
		"Second line",
		"Third line",
	}

	if err := WriteLines(filePath, want); err != nil {
		t.Fatalf("WriteLines() error = %v", err)
	}

	got, err := ReadLines(filePath)
	if err != nil {
		t.Fatalf("ReadLines() error = %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ReadLines() = %#v, want %#v", got, want)
	}
}

func TestReadLinesMissingFile(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "missing.txt"))

	lines, err := ReadLines(filePath)

	if err == nil {
		t.Fatal("ReadLines() error = nil, want error")
	}

	if lines != nil {
		t.Fatalf("ReadLines() = %#v, want nil", lines)
	}
}

func TestWriteJSONAndReadJSON(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "config.json"))

	type Config struct {
		App    string `json:"app"`
		Active bool   `json:"active"`
	}

	want := Config{
		App:    "FileUtility",
		Active: true,
	}

	if err := WriteJSON(filePath, want); err != nil {
		t.Fatalf("WriteJSON() error = %v", err)
	}

	var got Config

	if err := ReadJSON(filePath, &got); err != nil {
		t.Fatalf("ReadJSON() error = %v", err)
	}

	if got != want {
		t.Fatalf("ReadJSON() = %+v, want %+v", got, want)
	}
}

func TestReadJSONInvalidJSON(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "invalid.json"))

	if err := os.WriteFile(
		string(filePath),
		[]byte(`{"name":`),
		0644,
	); err != nil {
		t.Fatalf("failed to create invalid JSON file: %v", err)
	}

	var result map[string]any

	err := ReadJSON(filePath, &result)

	if err == nil {
		t.Fatal("ReadJSON() error = nil, want decoding error")
	}
}

func TestReadJSONMissingFile(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "missing.json"))

	var result map[string]any

	err := ReadJSON(filePath, &result)

	if err == nil {
		t.Fatal("ReadJSON() error = nil, want error")
	}
}

func TestWriteJSONInvalidValue(t *testing.T) {
	dir := t.TempDir()

	filePath := Filepath(filepath.Join(dir, "invalid.json"))

	invalidValue := make(chan int)

	err := WriteJSON(filePath, invalidValue)

	if err == nil {
		t.Fatal("WriteJSON() error = nil, want marshal error")
	}
}

func TestMustNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Must(nil) panicked unexpectedly: %v", r)
		}
	}()

	Must(nil)
}

func TestMustPanics(t *testing.T) {
	wantErr := errors.New("test error")

	defer func() {
		r := recover()

		if r == nil {
			t.Fatal("Must() did not panic")
		}

		gotErr, ok := r.(error)
		if !ok {
			t.Fatalf("Must() panic value = %T, want error", r)
		}

		if !errors.Is(gotErr, wantErr) {
			t.Fatalf("Must() panic = %v, want %v", gotErr, wantErr)
		}
	}()

	Must(wantErr)
}
