package models

type User struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
