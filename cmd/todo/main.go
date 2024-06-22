package main

import (
	"flag"
	"fmt"

	"github.com/dragon-huang0403/todo-go/pkg/config"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "", "config file path")
	flag.Parse()

	conf := config.Config{}.Default()
	conf.ConfigFile = configFile

	appConfig := Config{}
	err := config.GetConfig(conf, &appConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(appConfig)
}
