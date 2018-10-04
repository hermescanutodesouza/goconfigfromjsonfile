package goconfigfromjsonfile

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Config struct {
	Host     string
	Database string
	User     string
	Password string
	Server   string
}

func Getconfig() Config {
	var config Config
	err := gonfig.GetConf("config.json", &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
