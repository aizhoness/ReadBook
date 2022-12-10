package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	DBParam struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"db"`
}

func ProvideConfig() *Config {
	conf := &Config{}
	file, err := os.Open("config/base.yaml")
	if err != nil {
		//temporary we will panic errors
		panic(err)
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&conf); err != nil {
		//temporary we will panic errors
		panic(err)
	}

	return conf
}

var Options = ProvideConfig
