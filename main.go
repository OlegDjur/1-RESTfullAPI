package main

import (
	"net/http"
	"ta/web/users"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// db.Migrate()
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", users.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", users.GetUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", users.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}", users.DeleteUserHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
