package main

import (
	// Standard library packages
	"fmt"
	"net/http"

	"github.com/nischayjain4948/golang-webapis/controllers"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Instantiate a new router
	fmt.Println("This file is used to run the main.go file....")
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	// handle the routes..
	r.GET("/test", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Fire up the server
	http.ListenAndServe("localhost:5000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb+srv://Nischay:Nischay@crudd.qoju4.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	return s
}
