package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
}

type Users []User
