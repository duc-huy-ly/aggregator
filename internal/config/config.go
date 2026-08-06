package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/duc-huy-ly/aggregator/internal/database"
	"github.com/google/uuid"
)

const FILENAME string = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"Db_url,omitempty"`
	ConnectionString  string `json:"Connection_string,omitempty"`
	Current_user_name string `json:"Current_user_name,omitempty"`
}

func (c Config) DatabaseURL() string {
	if c.ConnectionString != "" {
		return c.ConnectionString
	}
	return c.Db_url
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

func (c *Config) RegisterUser(name string, db *database.Queries) (database.User, error) {
	newUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newUser, err := db.CreateUser(ctx, newUserParams)
	if err != nil {
		return database.User{}, fmt.Errorf("create user: %w", err)
	}
	return newUser, nil

}
