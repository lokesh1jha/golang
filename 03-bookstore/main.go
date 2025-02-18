package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lokesh1jha/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.ResgisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
