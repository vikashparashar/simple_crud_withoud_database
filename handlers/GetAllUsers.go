package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vikahparasha/simple_rest_without_database/models"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("request Method Not Allowed !")
	} else {
		fmt.Println("Total No Of Users : ", len(models.Users))
		json.NewEncoder(w).Encode(models.Users)
	}
}
