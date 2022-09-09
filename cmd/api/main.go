package main

import (
	"os"

	"github.com/ArdentK/db-cp-final/config"
	"github.com/ArdentK/db-cp-final/server"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// app, err := server.NewApp()
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }

	if err := server.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
