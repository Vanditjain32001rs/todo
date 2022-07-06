package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/helper"
	"todo/models"
)

func MarkTaskComplete(w http.ResponseWriter, r *http.Request) {

	var taskID models.UserTaskID
	msg := json.NewDecoder(r.Body).Decode(&taskID)
	if msg != nil {
		log.Printf("MarkTaskComplete : Error in decoding the json body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isCompleteErr := helper.MarkCompleteQuery(taskID.TaskID)

	if isCompleteErr != nil {
		log.Printf("MarkTaskCOmplete : Error in updating the task ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	taskDesc, err := helper.FetchTaskQuery(taskID.TaskID)
	if err != nil {
		log.Printf("MarkTaskComplete : Error in fetching the task description")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, jsonErr := json.Marshal(taskDesc)
	if jsonErr != nil {
		log.Printf("MarkTaskComplete : Error in converting to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, wErr := w.Write(jsonData)
	if wErr != nil {
		log.Printf("MarkTaskComplete : Error in writing the json data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
