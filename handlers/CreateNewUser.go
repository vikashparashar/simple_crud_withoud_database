package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vikahparasha/simple_rest_without_database/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var found bool
	if r.Method != "POST" {
		log.Println("request Method Not Allowed !")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		var newuser models.User
		json.NewDecoder(r.Body).Decode(&newuser)
		for _, v := range models.Users {
			if v.ID == newuser.ID {
				found = true
			}
		}

		if found == false {
			models.Users = append(models.Users, newuser)
			w.WriteHeader(http.StatusOK)
			fmt.Println("\t\tNew User Create Process Is Success ")
			json.NewEncoder(w).Encode(models.Users)
		} else {
			fmt.Println("Error : Id Already Exists")
			w.Write([]byte("Error : Id Already Exists"))
		}

	}
}
