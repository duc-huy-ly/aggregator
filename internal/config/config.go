package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const FILENAME string = ".gatorconfig.json"

type Config struct {
	Db_url            string
	Current_user_name string
}

func Read() Config {
	configFile, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("Error fetching config file : %v\n", err)
		return Config{}
	}
	file, err := os.Open(configFile)
	if err != nil {
		return Config{}
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	c := Config{}
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Printf("Error decoding the file")
		return c
	}
	return c
}

func (c *Config) SetUser(user_name string) error {
	if user_name == "" {
		return fmt.Errorf("Error, username required")
	}
	c.Current_user_name = user_name
	err := write(c)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	return nil

}

func write(c *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("Error MarshalIndent, %v\n", err)
	}
	os.WriteFile(path, data, 0644)
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/" + FILENAME, nil

}
