package config

import (
	"os"
	"encoding/json"
	"fmt"
)

func Read() (Config, error) {

	location, err := getGatorConfigLocation()
	if err != nil {
		return Config{}, err
	}

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
	// Setting the slice of content to read only bytes needed
	if err := json.Unmarshal(content[:n], &config); err != nil {
		return Config{}, fmt.Errorf("error in reading as json: %s\n", err)
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return c.writeToConfig()
}

func (c *Config) PrintConfig() {
	fmt.Printf("db_url: %s\ncurrent_user_name: %s\n", c.DBUrl, c.CurrentUserName)
}

// Internal function to write the json
func (c *Config) writeToConfig() error {

	location, err := getGatorConfigLocation()
	if err != nil {
		return err
	}

	// Marshal our config struct into json bytes
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	// Create overrwrites!
	file, err := os.Create(location)
	if err != nil {
		return fmt.Errorf("error in opening file %s: %s\n", location, err)
	}

	defer file.Close()

	// Write data into file location
	_, writeErr := file.Write(data)
	if writeErr != nil {
		return fmt.Errorf("error in writing to file %s: %s\n", location, writeErr)
	}

	return nil
}

// Internal method to return gator config location
func getGatorConfigLocation() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error in locating home directory: %s\n", err)
	}
	location := homeDir + "/" + gatorConfigName
	return location, nil
}