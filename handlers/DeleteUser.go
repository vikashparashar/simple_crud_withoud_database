package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vikahparasha/simple_rest_without_database/models"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// declearing varibales as per requirment
	var (
		id     string
		err    error
		found  = true
		index  int
		rest   []models.User
		pre    []models.User
		post   []models.User
		UserId int

		vars map[string]string
	)
	// set up content type header
	w.Header().Set("Content-Type", "application/json")
	// checking for request method

	if r.Method != "DELETE" {
		log.Println("Error : Request Method Not Allowed !")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		vars = mux.Vars(r)
		id = vars["id"]
		UserId, err = strconv.Atoi(id)

		if err != nil {
			log.Fatalln("Error : Failed To Convert String Type Id Into Int Type")
		}

		for i, v := range models.Users {
			if v.ID == UserId {
				found = true
				index = i
			}
		}

		if !found {
			w.Write([]byte("Error : No User Found With Given Id"))
			log.Println("Error : No User Found")
		} else {
			pre = models.Users[:index]
			post = models.Users[index+1:]
			rest = append(rest, pre...)
			rest = append(rest, post...)
			models.Users = rest
			w.WriteHeader(http.StatusOK)
			log.Println("Success : User Deleted !")
			json.NewEncoder(w).Encode(models.Users)
		}
	}
}
