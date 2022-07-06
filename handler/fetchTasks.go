package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/helper"
)

func FetchTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("UserID").(string)

	task, fetchTaskErr := helper.FetchAllTaskQuery(userID)
	if fetchTaskErr != nil {
		log.Printf("FetchTask : Error in fetching user tasks.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(task)
	if jsonErr != nil {
		log.Printf("FetchTask : Error in converting data to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("FetchTask : Error in writing json data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
