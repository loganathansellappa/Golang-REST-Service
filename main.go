package main

import (
	"FruitSale/app/config"
	"FruitSale/app/db"
	"github.com/rs/cors"
	"log"
	"net/http"
)

/*
main
	Starts the http sever with hardcoded port number
    Populates dummy data
*/

func main() {
	log.Println("Loading Seed data....")
	db.DB.Seed()
	handler := cors.AllowAll().Handler(config.InitRoutes())
	log.Println("App Started.....")
	http.ListenAndServe(":3000", handler)
}
