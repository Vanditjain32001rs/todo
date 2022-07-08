package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo/helper"
	"todo/models"
)

func FetchTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("UserID").(string)

	var page models.UserFetchTask
	PageID := r.URL.Query().Get("pageNo")
	page.PageNo, _ = strconv.Atoi(PageID)
	TaskLimit := r.URL.Query().Get("taskLimit")
	page.TaskSize, _ = strconv.Atoi(TaskLimit)

	if page.TaskSize == 0 {
		page.TaskSize = 5
	}
	task, fetchTaskErr := helper.FetchAllTaskQuery(userID, page.PageNo-1, page.TaskSize)
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
