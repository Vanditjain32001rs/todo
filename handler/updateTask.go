package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/helper"
	"todo/models"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var user models.UserTasks
	msg := json.NewDecoder(r.Body).Decode(&user)
	if msg != nil {
		log.Printf("UpdateTask : Error in reading the request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updateTaskErr := helper.UpdateTaskQuery(user.TaskID, user.TaskDescription)
	if updateTaskErr != nil {
		log.Printf("UpdateTask : Error in updating the task")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	taskString, fetchErr := helper.FetchTaskQuery(user.TaskID)
	if fetchErr != nil {
		log.Printf("UpdateTask : Error in fetching the updated task")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(taskString)
	if jsonErr != nil {
		log.Printf("UpdateTask : Error in converting to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("UpdateTask : Error in writing json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
