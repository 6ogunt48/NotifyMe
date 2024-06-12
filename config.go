package main

import (
	"github.com/BurntSushi/toml"
	"os"
)

func LoadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
