package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     uint16 `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	Server struct {
		Host string `json:"host"`
		Port uint16 `json:"port"`
	} `json:"server"`
	Properties `json:"properties"`
}

type Properties struct {
	Debug bool `json:"debug"`
}

func (config *Config) GetDatabaseConfigString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password,
		config.Database.DBName,
	)
}

func (config *Config) GetServerConfigString() string {
	return fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
}

func LoadConfigFile(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	config := new(Config)

	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, err
}
