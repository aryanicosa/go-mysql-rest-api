package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/users", controller.MiddlewareAuth(controller.GetAllUser(db)))
	router.HandleFunc("/user", controller.MiddlewareAuth(controller.CreateUser(db))).Methods("POST")
	router.HandleFunc("/user/{userId}", controller.MiddlewareAuth(controller.GetUser(db))).Methods("GET")
	router.HandleFunc("/user/{userId}", controller.MiddlewareAuth(controller.DeleteUser(db))).Methods("DELETE")
	router.HandleFunc("/user/{userId}", controller.MiddlewareAuth(controller.UpdateUser(db))).Methods("PUT")

	http.Handle("/", router)

	address := "localhost:9000"
	fmt.Printf("Server started at %s", address)

	http.ListenAndServe(address, router)
}
