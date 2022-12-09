package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ApplicationConfig struct {
	Address string
}

type Config struct {
	ApplicationConfig `yaml:"application"`
}

func ProvideConfig() *Config {
	conf := &Config{}
	data, err := ioutil.ReadFile("/base.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return conf
}

var Options = ProvideConfig
