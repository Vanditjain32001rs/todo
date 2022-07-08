package models

type UserTasks struct {
	TaskID          string `json:"TaskId" db:"task_id"`
	TaskDescription string `json:"TaskDesc" db:"task_description"`
}

type UserTaskDesc struct {
	TaskDescription string `json:"TaskDesc" db:"task_description"`
}

type UserTaskID struct {
	TaskID string `json:"TaskID" db:"task_id"`
}

type UserFetchTask struct {
	PageNo   int `json:"PageNo"`
	TaskSize int `json:"NoOfTask"`
}

type ContextMap struct {
	CtxMap map[string]string
}
