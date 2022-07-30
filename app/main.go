package main

import (
	"log"
	"os"

	"github.com/hmrbcnto/prescription-api/app/server"
	"github.com/hmrbcnto/prescription-api/config"
	db "github.com/hmrbcnto/prescription-api/infastructure/db/mongo"
)

func main() {

	// Getting uri string
	config, err := config.LoadConfig()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Creating mongodb connection
	client, err := db.NewConnection(config.DbConfig.MongoURI)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	srv := server.New(client)

	if err = srv.ListenAndServe(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
