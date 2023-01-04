package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vikahparasha/simple_rest_without_database/handlers"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", handlers.AllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.SingleUser).Methods("GET")
	router.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}/{name}/{age}", handlers.UpdateUserOne).Methods("PUT")
	router.HandleFunc("/api/users", handlers.UpdateUserTwo).Methods("PUT")
	router.HandleFunc("/api/users/{id}", handlers.DeleteUser).Methods("DELETE")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
