package backend

import (
	"flag"
	"log"

	"github.com/0mlml/cfgparser"
)

var config *cfgparser.Config

var configFile = flag.String("config", "server.cfg", "Path to the config file")

func ConfigInit() {
	defaultConfig := &cfgparser.Config{}
	defaultConfig.Literal(
		map[string]bool{},
		map[string]string{
			"db_path":        "tiramisu.db",
			"questions_path": "questions.json",
		},
		map[string]int{
			"port": 8080,
		},
		map[string]float64{},
	)

	cfgparser.SetDefaultConfig(defaultConfig)

	config = &cfgparser.Config{}
	if err := config.From(*configFile); err != nil {
		log.Fatalf("Error parsing config file: %v\nMake sure you have created %s from example.cfg", err, *configFile)
	}
}
