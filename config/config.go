package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type smtp struct {
	Username string `json:username`
	Password string `json:password`
	Host     string `json:host`
}

type Config struct {
	Smtp smtp `json:"smtp"`
}

func FromFile(filename string) (*Config, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, err
}
