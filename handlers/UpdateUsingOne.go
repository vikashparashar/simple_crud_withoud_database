package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vikahparasha/simple_rest_without_database/models"
)

func UpdateUserOne(w http.ResponseWriter, r *http.Request) {
	// 1st trick
	// declearing varibales as per requirment
	var (
		user          models.User
		found         bool = false
		index         int
		newUsersSlice []models.User
		name          string
		age           string
		newAge        int
		id            string
		newId         int
		err           error
		params        map[string]string
		a             []models.User
		b             []models.User
	)

	// set up content type header
	w.Header().Set("Content-Type", "application/json")
	// checking for request method
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Error : Request Method Not Allowed !")
	} else {
		params = mux.Vars(r)
		id = params["id"]
		newId, err = strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}
		name = params["name"]
		age = params["age"]
		newAge, err = strconv.Atoi(age)
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
			a = models.Users[:index]
			b = models.Users[index+1:]
			newUsersSlice = append(newUsersSlice, a...)
			newUsersSlice = append(newUsersSlice, b...)
			newUsersSlice = append(newUsersSlice, user)

			models.Users = newUsersSlice
			w.WriteHeader(http.StatusOK)
			log.Println("Success :  User Found With With Given Id !")
			log.Println("Success :  User Updated !")
			json.NewEncoder(w).Encode(models.Users)
		} else {
			log.Println("Error : No User Found With Given Id !")
			w.Write([]byte("Error : No User Found With Given Id !"))
		}
	}
}
