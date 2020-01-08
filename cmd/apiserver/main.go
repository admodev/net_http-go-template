package main

import (
	"flag"
	"log"

	"github.com/create-go-app/net_http-go-template/internal/apiserver"
)

var (
	configPath string
)

func init() {
	// Looking for flags
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yml", "path to config file")
}

func main() {
	// Parse flags
	flag.Parse()

	// Create new config
	config := apiserver.NewConfig(configPath)

	// Create new server
	server := apiserver.New(config)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
