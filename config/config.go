package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server string `json:"server"`
	Sender string `json:"defaultSender"`
}

func Parse(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	if err := json.Unmarshal(buf, config); err != nil {
		return nil, err
	}
	return config, nil
}
