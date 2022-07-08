package handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"todo/helper"
	"todo/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	msg := json.NewDecoder(r.Body).Decode(&newUser)
	if msg != nil {
		log.Printf("CreateUser : Error in decoding the json data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		log.Printf("CreateUser : Error in hashing the password.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ID, newUserErr := helper.CreateUser(newUser.Name, newUser.Email, newUser.Username, string(bytes))
	if newUserErr != nil {
		log.Printf("CreateUser : Error in returnig user id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//todo empty interface do read about it
	jsonData, jsonErr := json.Marshal(ID)
	if jsonErr != nil {
		log.Printf("CreateUser : Error in creating json file.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("Write Error : Error in writing json data.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
