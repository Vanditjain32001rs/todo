package models

type UserDetails struct {
	Name     string `json:"name" db:"name"`
	Username string `json:"userName" db:"username"`
}
