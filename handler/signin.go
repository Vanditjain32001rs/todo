package handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	"todo/helper"
	"todo/models"
)

func SignIn(w http.ResponseWriter, req *http.Request) {

	var cred models.Credential
	msgErr := json.NewDecoder(req.Body).Decode(&cred)
	if msgErr != nil {
		log.Printf("Signin : Error in decoding the json data.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pass, getPassErr := helper.GetPassword(cred.Username)
	if getPassErr != nil {
		log.Printf("Signin : Error in retreiving the password from database.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if compareErr := bcrypt.CompareHashAndPassword([]byte(pass), []byte(cred.Password)); compareErr != nil {
		log.Printf("Signin : Error in comparing the passwords.")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, getIdErr := helper.GetUserID(cred.Username)
	if getIdErr != nil {
		log.Printf("SignIn : Error in retreiving the User ID.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ExpiresAt := time.Now().Add(8 * time.Hour)
	var sess = &models.Session{}
	sess.ID = userID
	sess.Username = cred.Username
	sess.ExpiryTime = ExpiresAt

	sessID, getSessIdErr := helper.CreateSession(sess)
	if getSessIdErr != nil {
		log.Printf("SignIn : Error in retreiving the session id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var sessMap = make(map[string]string)
	sessMap["Session ID"] = sessID
	jsonData, jsonErr := json.Marshal(sessMap)
	if jsonErr != nil {
		log.Printf("SignIn : Error in converting to json data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("SignIn : Error in writing json data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
