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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var found = true
	var index int
	var rest []models.User
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		log.Println("request Method Not Allowed !")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		UserId, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalln("failed to convert string type id into int type")
		}

		for i, v := range models.Users {
			if v.ID == UserId {
				found = true
				index = i
			}
		}

		if found == false {
			w.Write([]byte("Error : No User Found With Given Id"))
			fmt.Println("Error : No User Found")
		} else {
			pre := models.Users[:index]
			post := models.Users[index+1:]
			rest = append(rest, pre...)
			rest = append(rest, post...)
			models.Users = rest
			w.WriteHeader(http.StatusOK)
			fmt.Println("\t\tDelete User Process Is Success  ")
			json.NewEncoder(w).Encode(models.Users)
		}
	}
}
