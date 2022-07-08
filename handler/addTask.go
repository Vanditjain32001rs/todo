package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"todo/helper"
	"todo/models"
	//"todo/models"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	//var userTask models.UserTaskDesc
	//msg := json.NewDecoder(r.Body).Decode(&userTask)
	//if msg != nil {
	//	log.Printf("AddTask : Error in decoding the json body")
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	userID := r.Context().Value("ID").(models.ContextMap)

	f, _, err := r.FormFile("excel")
	if err != nil {
		log.Printf("Error in formfile()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		log.Printf("error copying to bytes.buffer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileName := fmt.Sprintf("./%v", time.Now().Unix())
	create, CreateTempErr := os.Create(fileName)
	if CreateTempErr != nil {
		log.Printf("AddTask : Error in creating excel file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, CopyErr := io.Copy(create, buf)
	if CopyErr != nil {
		log.Printf("AddTask : Error in copying data to temp file")
		return
	}
	//fmt.Println(n)

	file, OpenErr := excelize.OpenFile(fileName)
	if OpenErr != nil {
		log.Printf("AddTask : Error in opening temp file")
		return
	}
	//taskIDarr := make([]string, 0, 1)
	rows, rowErr := file.GetRows("Sheet1")
	if rowErr != nil {
		log.Printf("AddTask : Error in getting rows in temp file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	taskMap := make(map[int]string)
	for n, row := range rows {
		for _, colCell := range row {
			taskMap[n] = colCell
		}
	}

	taskID, AddTaskErr := helper.AddTaskQuery(userID.CtxMap["UserID"], taskMap)
	if AddTaskErr != nil {
		log.Printf("AddTask : Error in Add task query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if CloseErr := file.Close(); CloseErr != nil {
		log.Printf("AddTask : Error in closing the temp file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	removeErr := os.Remove(fileName)
	if removeErr != nil {
		log.Printf("AddTask : Error in removing temp file")
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
