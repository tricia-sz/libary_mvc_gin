package models

import "time"

type User struct {
	ID        int64     `json:"ID"`
	Name      string    `json:"nome"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
