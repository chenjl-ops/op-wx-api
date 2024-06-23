package main

import (
	log "github.com/sirupsen/logrus"
	"op-wx-api/internal/app/server"
)

func main() {
	log.Info("Start OP-WX-API Service ....")
	err := server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
