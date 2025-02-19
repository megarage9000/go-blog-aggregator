package config

import (
	"os"
	"encoding/json"
	"fmt"
)

func Read() (Config, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("error in locating home directory: %s\n", err)
	}

	location := homeDir + "/" + gatorConfigName

	// Opening file
	file, err := os.Open(location)
	if err != nil {
		return Config{}, fmt.Errorf("error in opening file %s: %s\n", location, err)
	}

	defer file.Close()

	// Reading file 
	content := make([]byte, fileSize)

	n, err := file.Read(content)
	if err != nil {
		return Config{}, fmt.Errorf("error in reading file %s: %s\n", location, err)
	}

	if n == 0 {
		return Config{}, fmt.Errorf("file read 0 bytes\n")
	}

	var config Config
	if err := json.Unmarshal(content[:n], &config); err != nil {
		return Config{}, fmt.Errorf("error in reading as json: %s\n", err)
	}

	return config, nil
}

