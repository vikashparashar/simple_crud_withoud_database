package handlers

import (
	"encoding/json"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vikahparasha/simple_rest_without_database/models"
)

func SingleUser(w http.ResponseWriter, r *http.Request) {
	// set up content type header
	w.Header().Set("Content-Type", "application/json")

	// declearing varibales as per requirment
	var (
		found  bool
		user   models.User
		vars   map[string]string
		id     string
		UserId int
		err    error
	)
	// checking for request method
	if r.Method != "GET" {
		log.Println("Error : Request Method Not Allowed !")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		vars = mux.Vars(r)
		id = vars["id"]
		UserId, err = strconv.Atoi(id)

		if err != nil {
			log.Fatalln("Failed Error : can not convert string type id into int type")
		}
		for _, v := range models.Users {
			if v.ID == UserId {
				found = true
				user = v
			}

		}
		if !found {
			log.Println("Error : No User Found !")
			w.Write([]byte("Error : No User Found With Given Id"))
		} else {
			log.Println("Success : User Found !")
			json.NewEncoder(w).Encode(user)
		}
	}
}
