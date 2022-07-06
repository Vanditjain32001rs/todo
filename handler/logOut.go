package handler

import (
	"log"
	"net/http"
	"todo/helper"
)

func LogOut(w http.ResponseWriter, r *http.Request) {

	sessID, ok := r.Context().Value("SessionID").(string)
	if !ok {
		return
	}

	logOutErr := helper.DeleteSession(sessID)
	if logOutErr != nil {
		log.Printf("LogOut : Error in deleting the session")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
