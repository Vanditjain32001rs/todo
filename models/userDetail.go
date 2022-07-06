package models

type UserDetails struct {
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
}
