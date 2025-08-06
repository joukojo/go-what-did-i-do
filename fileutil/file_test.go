package fileutil_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/joukojo/go-what-did-i-do/fileutil"
)

func TestWriteFileAndReadFile(t *testing.T) {
	tmpFile := "testfile.json"
	defer func() {
		if err := os.Remove(tmpFile); err != nil {
			t.Errorf("Failed to remove temp file: %v", err)
		}
	}()

	type testStruct struct {
		Name string
		Age  int
	}
	input := testStruct{Name: "Alice", Age: 30}

	err := fileutil.WriteFile(tmpFile, input)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	data, err := fileutil.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	var output testStruct
	err = json.Unmarshal(data, &output)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(input, output) {
		t.Errorf("Expected %v, got %v", input, output)
	}
}

func TestReadFileNotExist(t *testing.T) {
	_, err := fileutil.ReadFile("nonexistentfile.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestWriteFileMarshalError(t *testing.T) {
	f, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() {
		if err := os.Remove(f.Name()); err != nil {
			t.Errorf("Failed to remove temp file: %v", err)
		}
	}()

	defer func() {
		err := f.Close()
		if err != nil {
			t.Errorf("Failed to close temp file: %v", err)
		}
	}()

	ch := make(chan int) // channels cannot be marshaled to JSON
	err = fileutil.WriteFile(f.Name(), ch)
	if err == nil {
		t.Error("Expected error when marshaling unsupported type, got nil")
	}
}
