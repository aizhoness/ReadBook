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
