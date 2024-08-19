package main

import (
	"log"
	"os"

	"github.com/backend/routes"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvs()
	if err != nil {
		log.Fatal(err)
	}

	f := newHTTP()

	db, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	routes.InitRoutes(f, db)

	log.Fatal(f.Listen(":" + os.Getenv("SERVER_PORT")))
}
