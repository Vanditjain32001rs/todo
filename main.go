package main

import (
	"log"
	"todo/database"
	"todo/server"
)

func main() {

	dbConnection := database.ConnectAndMigrate("localhost", "5433", "todo", "local", "local", "disable")
	if dbConnection != nil {
		panic(dbConnection)
	}

	log.Printf("Connected")

	srv := server.SetUpRoutes()
	start := srv.Run(":8080")
	if start != nil {
		log.Printf("main : Error in listenig to the requests.")
		log.Fatal(start)
	}

}
