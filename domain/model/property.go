package model

import (
	"github.com/google/uuid"
)

type Property struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	Length        float64   `json:"length"`
	Width         float64   `json:"width"`
	Area          float64   `json:"area"`
	Floor         float64   `json:"floor"`
	NumberOfFloor float64   `json:"number_of_floor"`
}

type Properties []Property
