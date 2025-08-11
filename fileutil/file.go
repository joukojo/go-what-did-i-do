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
	// #nosec
	err = os.WriteFile(filename, data, 0600) // write with read/write permissions for owner

	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

// ReadFile reads the contents of a file and returns it as a byte slice.
func ReadFile(filename string) ([]byte, error) {
	// #nosec
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

// GetDataDirectory retrieves the data directory for the application.
func GetDataDirectory() error {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return err
	}
	dataDir := fmt.Sprintf("%s/%s", homeDir, ".what-did-i-do")
	return EnsureDir(dataDir)
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

// EnsureDir checks if a directory exists and creates it if it does not.
func EnsureDir(path string) error {
	// Check if the directory exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Directory does not exist, so create it
		err = os.MkdirAll(path, 0750) // create with read/write/execute permissions for owner, and read/execute for group
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		fmt.Println("Directory created:", path)
	} else if err != nil {
		return fmt.Errorf("failed to check directory: %w", err)
	} else {
		fmt.Println("Directory already exists:", path)
	}
	return nil
}
