package controller

import (
	"database/sql"
	"encoding/json"
	"go-rest-api/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const USERNAME = "john"
const PASSWORD = "doe"

func MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		isValid := (username == USERNAME && password == PASSWORD)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("wrong username/password"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetAllUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.Query("SELECT * FROM user")
		if err != nil {
			panic(err)
		}

		data := []model.User{}

		for users.Next() {
			var user model.User

			err = users.Scan(&user.ID, &user.Name, &user.Email, &user.Address)
			if err != nil {
				panic(err)
			}

			data = append(data, user)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := model.User{}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["userId"])

		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		err = db.QueryRow("SELECT id, name, email, address FROM user WHERE id = ?", id).
			Scan(&data.ID, &data.Name, &data.Email, &data.Address)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.User{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err := db.Exec("INSERT INTO user VALUES (?, ?, ?, ?)", "", data.Name, data.Email, data.Address)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.User{}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["userId"])

		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		err = db.QueryRow("SELECT id, name, email, address FROM user WHERE id = ?", id).
			Scan(&data.ID, &data.Name, &data.Email, &data.Address)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		_, err = db.Exec("DELETE FROM user WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Failed to delete", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.User{}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["userId"])

		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("UPDATE user SET name = ?, email = ?, address = ? WHERE id = ?",
			data.Name, data.Email, data.Address, id)

		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
