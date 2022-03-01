# Building GO-MySql CRUD rest api

This is a learning project to create CRUD REST API without using any framework.

In this project implement Basic Auth, so if we want to create any request we should define username and password (Encode64) in the request header : `Authorization: Basic Base64(username:password)`

Pre-requisite :
1. Install MySQL server on machine
2. Install go sql driver package        : `go get -u github.com/go-sql-driver/mysql`
3. Install Gorilla Mux route package    : `go get -u github.com/gorilla/mux`