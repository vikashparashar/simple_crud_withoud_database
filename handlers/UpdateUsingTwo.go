package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vikahparasha/simple_rest_without_database/models"
)

func UpdateUserTwo(w http.ResponseWriter, r *http.Request) {
	// 2nd trick
	// declearing varibales as per requirment
	var (
		requestBodyUser models.User
		newUserToAdd    models.User
		found           bool
		index           int
		newUsersSlice   []models.User
		newId           int
		err             error
		data            []byte
		a               []models.User
		b               []models.User
	)
	// set up content type header
	w.Header().Set("Content-Type", "application/json")
	// checking for request method
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Error : Request Method Not Allowed !")
	} else {

		// rading data from request body
		data, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("[ Error : Failed To Read Data From Request Body ]")
		}

		// unmarshal json data into models.User Struct
		err = json.Unmarshal(data, &requestBodyUser)
		if err != nil {
			log.Fatalln("[ Error : Failed To Unmarhsal The Request Body Data ]")
		}

		for i, v := range models.Users {
			if v.ID == newId {
				found = true
				index = i
			}
		}
		if !found {

			// assigning value to a newUserToAdd variable to add into slice
			newUserToAdd.ID = requestBodyUser.ID
			newUserToAdd.Name = requestBodyUser.Name
			newUserToAdd.Age = requestBodyUser.Age
			a = models.Users[:index]
			b = models.Users[index+1:]
			newUsersSlice = append(newUsersSlice, a...)
			newUsersSlice = append(newUsersSlice, b...)
			newUsersSlice = append(newUsersSlice, newUserToAdd)

			models.Users = newUsersSlice
			w.WriteHeader(http.StatusOK)
			log.Println("[ Success :  User Found With With Given Id ]")
			log.Println("[ Success :  User Updated ]")
			json.NewEncoder(w).Encode(models.Users)
		} else {
			log.Println("[ Error : No User Found With Given Id ]")
			w.Write([]byte("[ Error : No User Found With Given Id ]"))
		}
	}
}
