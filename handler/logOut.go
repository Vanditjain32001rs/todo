package handler

import (
	"log"
	"net/http"
	"todo/helper"
	"todo/models"
)

func LogOut(w http.ResponseWriter, r *http.Request) {

	//var ss sessUser
	ss := r.Context().Value("ID").(models.ContextMap)
	logOutErr := helper.DeleteSession(ss.CtxMap["SessID"])
	if logOutErr != nil {
		log.Printf("LogOut : Error in deleting the session")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
