package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nischayjain4948/golang-mysql-bookstore/pkg/routes"
)

func main() {
	fmt.Println("This is a project for creating the golang-mysql-bookstore ")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5000", r))

}
