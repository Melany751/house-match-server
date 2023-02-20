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
	PersonID uuid.UUID `json:"person_id"`
}

type Users []User

type UserOutput struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Theme    string    `json:"theme"`
	Person   Person    `json:"person"`
}

type UsersOutput []UserOutput
