package helper

import (
	"time"
	"todo/database"
	"todo/models"
)

func CreateUser(userName, userEmail, Username, Password string) (string, error) {

	createQuery := `INSERT INTO todo(name, email, username, password_hash)
					VALUES ($1, $2, $3, $4)
					RETURNING id`
	var userId string
	createErr := database.Data.Get(&userId, createQuery, userName, userEmail, Username, Password)

	return userId, createErr
}

func GetPassword(username string) (string, error) {

	query := `SELECT password_hash
			  FROM todo
			  WHERE username = $1`

	var hashPass string
	getPassErr := database.Data.Get(&hashPass, query, username)

	return hashPass, getPassErr
}

func GetID(sessID string) (string, error) {

	query := `SELECT id 
			  FROM sessions 
			  WHERE session_id=$1`

	var id string
	getIdErr := database.Data.Get(&id, query, sessID)

	return id, getIdErr
}

func GetUserID(username string) (string, error) {

	query := `SELECT id 
			  FROM todo 
			  WHERE username=$1`

	var id string
	getIdErr := database.Data.Get(&id, query, username)

	return id, getIdErr
}

func CreateSession(userSession *models.Session) (string, error) {
	query := `INSERT INTO sessions(id, expires_at) VALUES ($1,$2) returning session_id`

	var sessID string
	getSessionIdErr := database.Data.Get(&sessID, query, userSession.ID, userSession.ExpiryTime)

	return sessID, getSessionIdErr
}

func AddTaskQuery(userID, taskDesc string) (string, error) {

	query := `INSERT INTO tasks(id,task_description) 
		      VALUES($1,$2) 
			  RETURNING task_id`

	var taskID string
	addTaskErr := database.Data.Get(&taskID, query, userID, taskDesc)

	return taskID, addTaskErr
}

func UpdateTaskQuery(taskID, taskDesc string) error {

	query := `UPDATE tasks SET task_description=$1 WHERE task_id=$2`
	_, updateErr := database.Data.Exec(query, taskDesc, taskID)

	return updateErr
}

func FetchAllTaskQuery(userID string) (*[]models.UserTaskDesc, error) {

	query := `SELECT task_description
			  FROM tasks
			  WHERE id=$1 AND archived_at is null`

	var userTask []models.UserTaskDesc
	fetchTaskErr := database.Data.Select(&userTask, query, userID)
	if fetchTaskErr != nil {
		return &userTask, fetchTaskErr
	}
	return &userTask, nil
}

func FetchTaskQuery(taskID string) (string, error) {
	query := `SELECT task_description
              FROM tasks
			  WHERE task_id=$1`
	var taskDesc string
	fetchErr := database.Data.Get(&taskDesc, query, taskID)
	if fetchErr != nil {
		return "", fetchErr
	}

	return taskDesc, nil
}

func DeleteTask(taskID string) error {

	currTime := time.Now()
	query := `UPDATE tasks 
              SET archived_at=$1 
			  WHERE task_id=$2`
	_, err := database.Data.Exec(query, currTime, taskID)

	return err
}

func DeleteSession(sessID string) error {
	currTime := time.Now()
	query := `UPDATE sessions
			  SET archived_at=$1
			  WHERE session_id=$2`
	_, err := database.Data.Exec(query, currTime, sessID)

	return err
}
func MarkCompleteQuery(taskID string) error {
	isComplete := true
	query := `UPDATE tasks 
			   SET is_complete =$1
               WHERE task_id=$2`
	_, err := database.Data.Exec(query, isComplete, taskID)

	return err
}

func SessionExist(sessID string) (bool, error) {

	var isExpired bool
	query := `SELECT count(*) > 0
			  FROM sessions 
			WHERE session_id=$1 AND archived_at is null and expires_at is null`
	checkSessErr := database.Data.Select(&isExpired, query, sessID)
	if checkSessErr != nil {
		return isExpired, checkSessErr
	}
	return isExpired, nil
}

func RefreshSessToken(sessID string) error {
	expiryTime := time.Now().Add(10 * time.Minute)
	updateQuery := `UPDATE sessions
					SET expires_at=$1
					WHERE session_id=$2`
	_, updateErr := database.Data.Exec(updateQuery, expiryTime, sessID)

	return updateErr
}

func GetUserDetails(userID string) (*models.UserDetails, error) {

	query := `SELECT name,username FROM todo WHERE id=$1`
	var userDetail = models.UserDetails{}
	GetUserDetailErr := database.Data.Get(&userDetail, query, userID)

	return &userDetail, GetUserDetailErr
}
