package server

import (
	"fmt"

	"amald/config"
	"amald/loader"
)

func init() {

}

/*
Main executes the amald server
*/
func Main() {
	// parse options
	args := parseArgs()

	// get config
	conf := config.Load(args.configPath)

	// load urls
	urls := loader.Load(conf.Loaders)

	fmt.Println(urls)
}
