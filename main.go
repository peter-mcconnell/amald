package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	config "github.com/pemcconnell/amald/config"
	"github.com/pemcconnell/amald/loaders"
	"github.com/pemcconnell/amald/notifiers"
	"github.com/pemcconnell/amald/storage"
)

const (
	VERSION string = "0.0.3"
)

var (
	configpath = flag.String("c", "./config.yaml",
		"[config] set the path for the yaml config file. This defautls to "+
			"./config.yaml")
	templatepath = flag.String("t", "reports/tmpl/",
		"[templates directory] set the path for the templates directory")
	loglevel = flag.String("v", "info",
		"[loglevel] set the verbosity of the log levels. Can be: debug, "+
			"info, warn, error, panic, fatal")
	reportonly = flag.Bool("r", false, "[report only] if true amald will not run a scan but will instead only display a report")
)

func init() {
	// parse flags
	flag.Parse()

	// Set logrus level
	if level, err := log.ParseLevel(*loglevel); err == nil {
		log.SetLevel(level)
	}
}

func main() {

	welcome()

	// load the config
	config, err := config.LoadConfig(*configpath)
	if err != nil {
		log.Fatalf("There was a problem setting the config")
	}

	// grab the loaders
	err = loaders.GetLoaders(config.ActiveLoaders)
	if err != nil {
		log.Fatalf("Failed to grab loaders")
	}

	// Load Data
	if _, err := storage.LoadData(config.Storage); err != nil {
		log.Warnf("Failed to LoadData")
	}

	// Store Data
	if !*reportonly {
		urls := loaders.CollectUrls()
		if _, err := storage.StoreData(*reportonly, urls, config.Storage); err != nil {
			log.Fatalf("Failed to StoreData")
		}
	}

	// call the notifiers to display/send the messages
	notifiers.FireNotifiers(*templatepath, storage.ExistingDataBytes, config.Reports)
}

// welcome displays the version number to the user
func welcome() {
	log.Info("\n---------------------------------\n"+
		" ~ amald v"+VERSION, " ~\n"+
		"---------------------------------")
}
