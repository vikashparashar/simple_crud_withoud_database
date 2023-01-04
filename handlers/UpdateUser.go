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

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// 1st trick

	var user models.User
	var found bool
	var index int
	var NewUsersSLlice []models.User
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("\t\tRequest Method Not Allowed !")
	} else {
		params := mux.Vars(r)
		id := params["id"]
		newId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}
		name := params["name"]
		age := params["age"]
		newAge, err := strconv.Atoi(age)
		if err != nil {
			log.Fatalln(err)
		}
		for i, v := range models.Users {
			if v.ID == newId {
				found = true
				index = i
			}

		}

		if found == true {
			user.ID = newId
			user.Name = name
			user.Age = newAge
			a := models.Users[:index]
			b := models.Users[index+1:]
			NewUsersSLlice = append(NewUsersSLlice, a...)
			NewUsersSLlice = append(NewUsersSLlice, b...)
			NewUsersSLlice = append(NewUsersSLlice, user)

			models.Users = NewUsersSLlice
			w.WriteHeader(http.StatusOK)
			fmt.Println("\t\t User Updated Successfully ")
			json.NewEncoder(w).Encode(models.Users)
		} else {
			fmt.Println("\t\tError : No User Found With Given Id")
			w.Write([]byte("Error : No User Found With Given Id"))
		}

	}

	// 2nd trick
	// var user models.User
	// w.Header().Set("Content-Type", "application/json")
	// if r.Method != "PUT" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// } else {
	// 	params := mux.Vars(r)
	// 	id := params["id"]

	// 	newId, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
}
