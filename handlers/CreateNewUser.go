package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vikahparasha/simple_rest_without_database/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// declearing varibales as per requirment
	var (
		found   bool
		newuser models.User
	)
	// checking for request method
	if r.Method != "POST" {
		log.Println("Error : Request Method Not Allowed !")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {

		json.NewDecoder(r.Body).Decode(&newuser)
		for _, v := range models.Users {
			if v.ID == newuser.ID {
				found = true
			}
		}

		if found == false {
			models.Users = append(models.Users, newuser)
			w.WriteHeader(http.StatusOK)
			log.Println("Success : New User Create !")
			json.NewEncoder(w).Encode(models.Users)
		} else {
			log.Println("Error : User With Given Id Is Already Exists !")
			w.Write([]byte("Error :User With Given Id Is Already Exists !"))
		}

	}
}
