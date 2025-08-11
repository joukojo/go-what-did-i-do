// Package fileutil provides utility functions for working with files.
package fileutil

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteFile writes data to a file.
func WriteFile(filename string, content any) error {
	data, err := json.MarshalIndent(content, "", "  ") // pretty format
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err

	}
	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

// ReadFile reads the contents of a file and returns it as a byte slice.
func ReadFile(filename string) ([]byte, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return data, err

}

// getUserHomeDir retrieves the user's home directory.
func getUserHomeDir() (string, error) {
	if os.Getenv("WDID_DEBUG") != "" {
		return "/tmp", nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

// GetDataFile retrieves data from a JSON file in the user's home directory under .what-did-i-do.
func GetDataFile(datafilename string) ([]byte, error) {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return nil, err
	}
	dataFilePath := fmt.Sprintf("%s/%s/%s", homeDir, ".what-did-i-do", datafilename)
	return ReadFile(dataFilePath)
}

// WriteDataFile writes data to a JSON file in the user's home directory under .what-did-i-do.
func WriteDataFile(datafilename string, content any) error {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return err
	}
	dataFilePath := fmt.Sprintf("%s/%s/%s", homeDir, ".what-did-i-do", datafilename)
	fmt.Println("Writing data file to:", dataFilePath)
	return WriteFile(dataFilePath, content)
}
