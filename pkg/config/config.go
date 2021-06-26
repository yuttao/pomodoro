package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port int `json:"port"`
}

func ReadConfig() (Config, error) {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalln("Config file is not found: ./config/config.json")
		return Config{Port: 8080}, err
	}

	var config Config
	json.Unmarshal(data, &config)
	return config, nil
}
