package fileutil

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func ReadFile(filename string) ([]byte, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return data, err

}

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

func GetDataFile(datafilename string) ([]byte, error) {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return nil, err
	}
	dataFilePath := fmt.Sprintf("%s/%s/%s", homeDir, ".what-did-i-do", datafilename)
	return ReadFile(dataFilePath)
}

func WriteDataFile(datafilename string, content []byte) error {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return err
	}
	dataFilePath := fmt.Sprintf("%s/%s/%s", homeDir, ".what-did-i-do", datafilename)
	return WriteFile(dataFilePath, content)
}
