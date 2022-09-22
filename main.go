package main

import (
	"FruitSale/app/config"
	"FruitSale/app/db"
	"log"
	"net/http"
)

/*
main
	Starts the http sever with hardcoded port number
    Populates dummy data
*/

func main() {
	log.Print("Listening...")
	db.DB.Seed()
	http.ListenAndServe(":3000", config.InitRoutes())
}
