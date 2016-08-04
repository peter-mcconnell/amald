package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Conf is a global that stores the Config
var Conf Config

/*
Load takes a path string to a yaml file and loads it's values to the
global &Conf (Config)
*/
func Load(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load config file at %s: %s", path, err)
	}

	// parse reads data and unmarshals it to the global &conf
	if err = parse(data); err != nil {
		log.Fatalf("Failed to parse config file data: %s", err)
	}

	return &Conf
}

func parse(data []byte) error {
	if err := yaml.Unmarshal(data, &Conf); err != nil {
		return err
	}
	return nil
}
