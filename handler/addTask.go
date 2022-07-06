package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/helper"
	"todo/models"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	var userTask models.UserTaskDesc
	msg := json.NewDecoder(r.Body).Decode(&userTask)
	if msg != nil {
		log.Printf("AddTask : Error in decoding the json body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID := r.Context().Value("UserID").(string)
	taskID, AddTaskErr := helper.AddTaskQuery(userID, userTask.TaskDescription)
	if AddTaskErr != nil {
		log.Printf("AddTask : Error in Add task query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(taskID)
	if jsonErr != nil {
		log.Printf("AddTask : Error in converting to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("AddTask : Error in writing the json data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
