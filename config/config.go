package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadConfig returns a Config struct. It builds the config using the provided
// filepath. loadDefaults called after the yaml file is scanned to assign
// defaults if they are needed.
func LoadConfig(path string) (defs.Config, error) {

	var (
		config defs.Config
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
func loadDefaults(config defs.Config) (defs.Config, error) {
	return config, nil
}
