package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vikahparasha/simple_rest_without_database/models"
)

func SingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var found bool
	var user models.User
	if r.Method != "GET" {
		log.Println("request Method Not Allowed !")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		UserId, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalln("\t\tfailed to convert string type id into int type")
		}
		for _, v := range models.Users {
			if v.ID == UserId {
				found = true
				user = v
			}

		}
		if found != true {
			log.Println("\t\tNo User Found With Given Id")
			w.Write([]byte("Error : No User Found With Given Id"))
		} else {
			fmt.Println("\t\tUser Found With Given Id")
			json.NewEncoder(w).Encode(user)
		}
	}
}
