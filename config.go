package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct{
	Server struct{
		Host string `yaml:"host"`
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct{
		Postgres struct{
			Mode bool `yaml:"mode"`
			Host string `yaml:"host"`
			Port int `yaml:"port"`
			DbName string `yaml:"dbname"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			SslMode string `yaml:"sslmode"`
		} `yaml:"postgres"`
	} `yaml:"database"`
}

func readConfigurationFile(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		logrus.WithField("error", err).Fatal("Failed to open configuration file.")
		return nil
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		logrus.WithField("error", err).Fatal("Failed to parse configuration file.")
		return nil
	}

	return &cfg
}