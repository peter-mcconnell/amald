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
	VERSION string = "0.0.1"
)

var (
	configpath = flag.String("c", "./config.yaml",
		"[config] set the path for the yaml config file. This defautls to "+
			"./config.yaml")
	loglevel = flag.String("v", "info",
		"[loglevel] set the verbosity of the log levels. Can be: debug, "+
			"info, warn, panic, fatal")
)

func init() {
	// parse flags
	flag.Parse()

	// Set logrus level
	switch *loglevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	}
}

func main() {

	welcome()

	// load the config
	config, err := config.LoadConfig(*configpath)
	exitOnError(err, "There was a problem setting the config")

	// grab the loaders
	err = loaders.GetLoaders(config.ActiveLoaders)
	exitOnError(err, "Failed to grab loaders")

	// collect up some urls from the loaders
	urls := loaders.CollectUrls()
	exitOnError(err, "Failed to exec loaders")

	// storage
	err = storage.StoreData(urls, config.Storage)
	exitOnError(err, "Failed to StoreData")

	// call the notifiers to display/send the messages
	err = notifiers.FireNotifiers(urls, config.Reports)
	exitOnError(err, "Failed to FireNotifiers()")
}

// exitOnError checks that and error is not nil. If the passed value is an
// error, it is logged and the program exits with an error code of 1
func exitOnError(err error, prefix string) {
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal(prefix)
	}
}

// welcome displays the version number to the user
func welcome() {
	log.Info("\n---------------------------------\n"+
		" ~ amald v"+VERSION, " ~\n"+
		"---------------------------------")
}
