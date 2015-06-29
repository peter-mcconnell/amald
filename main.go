package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/config"
	//"github.com/pemcconnell/amald/defs"
	"github.com/pemcconnell/amald/loaders"
	"github.com/pemcconnell/amald/storage"
	"github.com/pemcconnell/amald/urltest"
	//	"github.com/pemcconnell/amald/notifiers"
)

const (
	VERSION string = "0.1.0"
)

var (
	configPath = flag.String("c", "./config.yaml",
		"[config] set the path for the yaml config file. This defaults to "+
			"./config.yaml")
	templatePath = flag.String("t", "reports/tmpl/",
		"[templates directory] set the path for the templates directory")
	logLevel = flag.String("v", "info",
		"[loglevel] set the verbosity of the log levels. Can be: debug, "+
			"info, warn, error, panic, fatal")
)

func init() {
	// parse flags
	flag.Parse()

	// Set logrus level
	if level, err := log.ParseLevel(*logLevel); err == nil {
		log.SetLevel(level)
	}
}

func main() {

	// load the config
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load the config from %s", *configPath)
	}

	// collect list of URLs from all possible loaders
	loaders.GetLoaders(cfg.Loaders)
	urls, err := loaders.CollectUrls()
	if err != nil {
		log.Fatalf("Failed to CollectUrls(): %s", err)
	}

	// test all of the urls
	scanResults, err := urltest.Batch(urls)
	if err != nil {
		log.Fatalf("Failed to Batch urltest: %s", err)
	}
	// if there aren't any urls found, don't continue
	if len(scanResults) == 0 {
		log.Fatal("No URLs found in loaders")
	}

	// grab a summary (compare current scan against old data)
	if cfg.Tests["storage"] {
		olddata, err := storage.LoadSiteDefsFromStorage(cfg.Storage["json"]["path"])
		if err != nil {
			log.Fatalf("Failed to get summary: %s", err)
		}
		// store latest test
		if cfg.Tests["storage"] {
			// store results
			merged := storage.MergeData(scanResults, olddata)
			storage.StoreScan(cfg.Storage["json"]["path"], merged)
		}
	}

	// fire off each notifier

}
