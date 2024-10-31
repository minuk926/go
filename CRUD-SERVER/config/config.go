package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	//func NewConfig(filePath string) (*Config, error) {
	var c = new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
		//return nil, err
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
		//return nil, err
	} else {
		return c
		//return c, nil
	}
}
