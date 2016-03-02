package main

import (
	"flag"

	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/logs"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/server"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/load-config"
)

var (
	flConfigFile = flag.String("config-file", "config-default.json", "Load config from file")

	log = logs.New("server")
)

func main() {
	flag.Parse()

	var cfg server.Config

	err := loadConfig.FromFileAndEnv(&cfg, *flConfigFile)

	if err != nil {
		log.Fatalln("Error loading config: ", err)
	}

	server.Start(cfg)
}
