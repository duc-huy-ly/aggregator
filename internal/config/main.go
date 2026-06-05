package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string
	Current_user_name string
}

func Read() Config {
	home, _ := os.UserHomeDir()
	fmt.Printf("%v\n", home)
	configFile := home + "/.gatorconfig.json"
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
