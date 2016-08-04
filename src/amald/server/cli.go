package server

import "flag"

/*
Args is a struct of arguments provided at run time
*/
type Args struct {
	configPath string
}

func parseArgs() *Args {
	configPath := flag.String("configPath", "./config.yaml", "A relative path to the yaml config file")
	flag.Parse()

	return &Args{
		configPath: *configPath,
	}
}
