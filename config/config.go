package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Port int `json:"port"`

	TokenSecret                string `json:"token_secret"`
	TokenExpiryInterval        int64  `json:"token_expiry_interval"`
	RefreshTokenExpiryInterval int64  `json:"refresh_token_expiry_interval"`

	DbName         string `json:"db_name"`
	DbHost         string `json:"db_host"`
	DbPort         int    `json:"db_port"`
	DbUser         string `json:"db_user"`
	DbPassword     string `json:"db_password"`
	DbMaxOpenConss int    `json:"db_max_open_conns"`
	DbMaxIdleConss int    `json:"db_max_idle_conns"`
}

func GetConfigFromFile(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
