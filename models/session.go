package models

import "time"

type Session struct {
	ID         string    `json:"id" db:"id"`
	Username   string    `json:"username" db:"username"`
	ExpiryTime time.Time `json:"expiry_time" db:"ExpiryTime"`
}
