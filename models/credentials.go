package models

type Credential struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password"`
}
