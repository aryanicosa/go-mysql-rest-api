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

	router.HandleFunc("/user/create", controller.MiddlewareBasicAuth(controller.CreateUser(db))).Methods("POST")
	router.HandleFunc("/user/login", controller.MiddlewareBasicAuth(controller.Login(db))).Methods("POST")
	router.HandleFunc("/users", controller.MiddlewareBearerAuth(controller.GetAllUser(db))).Methods("GET")
	router.HandleFunc("/user/{userId}", controller.MiddlewareBasicAuth(controller.GetUser(db))).Methods("GET")
	router.HandleFunc("/user/{userId}", controller.MiddlewareBasicAuth(controller.DeleteUser(db))).Methods("DELETE")
	router.HandleFunc("/user/{userId}", controller.MiddlewareBasicAuth(controller.UpdateUser(db))).Methods("PUT")

	http.Handle("/", router)

	address := "localhost:9000"
	fmt.Printf("Server started at %s", address)

	http.ListenAndServe(address, router)
}
