package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/helper"
	"todo/models"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	var taskID models.UserTaskID
	msg := json.NewDecoder(r.Body).Decode(&taskID)
	if msg != nil {
		log.Printf("DeleteTask : Error in deconding the json body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	delErr := helper.DeleteTask(taskID.TaskID)
	if delErr != nil {
		log.Printf("DeleteTask : Error in archiving the task")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	delString := fmt.Sprintf("Task Deleted")

	jsonData, jsonErr := json.Marshal(delString)
	if jsonErr != nil {
		log.Printf("DeleteTask : Error in converting to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("DeleteTask : Error in writing json body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
