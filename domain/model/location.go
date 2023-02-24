package model

import "github.com/google/uuid"

type Location struct {
	ID       uuid.UUID `json:"id"`
	Country  string    `json:"country"`
	City     string    `json:"city"`
	Province string    `json:"province"`
	District string    `json:"district"`
	Address  string    `json:"address"`
}

type Locations []Location
