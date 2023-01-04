package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vikahparasha/simple_rest_without_database/models"
)

// declearing varibales as per requirment

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	// set up content type header
	w.Header().Set("Content-Type", "application/json")
	// checking for request method
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Error : Request Method Not Allowed !")
	} else {
		log.Println("Success : Total No Of Users -  ", len(models.Users))
		err = json.NewEncoder(w).Encode(models.Users)
		if err != nil {
			log.Fatalln("Error : JSON Encoder Is Failed To Convert Models.Users Into JSON")
		}
	}
}
