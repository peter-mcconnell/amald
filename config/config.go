package config

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveLoaders map[string]map[string]string `json:activeloaders,omitempty`
	Reports       map[string]map[string]string `json:reports,omitempty`
}

// LoadConfig returns a Config struct. It builds the config using the provided
// filepath. loadDefaults called after the yaml file is scanned to assign
// defaults if they are needed.
func LoadConfig(path string) (Config, error) {

	var (
		config Config
	)

	// does file exist?
	path, _ = filepath.Abs(path)
	log.Debugf("LoadConfig %s", path)
	_, err := os.Stat(path)
	if err != nil {
		log.Warnf("config file not found: %s", path)
	} else {
		// read file
		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return config, err
		}
		// unmarshal
		if err = yaml.Unmarshal(yamlFile, &config); err != nil {
			return config, err
		}
	}

	// merge defaults
	config, err = loadDefaults(config)

	return config, nil
}

// loadDefaults is a placeholder at the moment. If default values are to be set
// to cover missing fields from the config yaml then they should be set here.
func loadDefaults(config Config) (Config, error) {
	return config, nil
}
