# Building GO-MySql CRUD rest api

This is a learning project to create CRUD REST API without using any framework.

In this project implement Basic Auth and Bearer Auth. 
- if we want to create any request to endpoint that implemented basic auth we should define username and password (Encode64) in the request header : `Authorization: Basic Base64(username:password)`
- if we want to create any request to endpoint that implemented bearer auth we should define token the request header : `Authorization: Bearer <token>`

Pre-requisite :
1. Install MySQL server on machine
2. Install go sql driver package        : `go get -u github.com/go-sql-driver/mysql`
3. Install Gorilla Mux route package    : `go get -u github.com/gorilla/mux`
4. Install JWT-Go                       : `go get github.com/dgrijalva/jwt-go`
5. Install bcrypt                       : `go get golang.org/x/crypto/bcrypt`

Demo (using PostMan) :
1. Create new user - basic auth - Created
![alt text](https://github.com/aryanicosa/go-mysql-rest-api/blob/main/doc/user_created.PNG)

2. Login and generate auth token - basic auth - OK
![alt text](https://github.com/aryanicosa/go-mysql-rest-api/blob/main/doc/user_login.PNG)

2. Get All User records - bearer auth - OK
![alt text](https://github.com/aryanicosa/go-mysql-rest-api/blob/main/doc/Correct_Bearer_Auth.PNG)

3. Get a User records - basic auth - OK
![alt text](https://github.com/aryanicosa/go-mysql-rest-api/blob/main/doc/Correct_basic_auth.PNG)