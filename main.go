package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/config"
	"github.com/pemcconnell/amald/loaders"
	//	"github.com/pemcconnell/amald/notifiers"
	//	"github.com/pemcconnell/amald/storage"
)

const (
	VERSION string = "0.1.0"
)

var (
	configpath = flag.String("c", "./config.yaml",
		"[config] set the path for the yaml config file. This defaults to "+
			"./config.yaml")
	templatepath = flag.String("t", "reports/tmpl/",
		"[templates directory] set the path for the templates directory")
	loglevel = flag.String("v", "info",
		"[loglevel] set the verbosity of the log levels. Can be: debug, "+
			"info, warn, error, panic, fatal")
	reportonly = flag.Bool("r", false, "[report only] if true amald will "+
		"not run a scan but will instead only display a report")
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

	// load the config
	cfg, err := config.Load(*configpath)
	if err != nil {
		log.Fatalf("Failed to load the config from %s", *configpath)
	}

	// collect list of URLs from all possible loaders
	loaders.GetLoaders(cfg.Loaders)
	loaders.CollectUrls()

	// test all of the urls

	// interact with storage
	//// store latest scan results

	//// compare latest scan with previously stored results

	// fire off each notifier

}
