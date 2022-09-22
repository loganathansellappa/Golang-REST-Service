package config

import (
	"FruitSale/app/controllers"
	"net/http"
)

/*
Request Matcher
*/
func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.HandleProductRequests)
	return mux
}
