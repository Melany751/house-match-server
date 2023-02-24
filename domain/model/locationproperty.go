package model

import "github.com/google/uuid"

type LocationProperty struct {
	ID         uuid.UUID `json:"id"`
	Lat        string    `json:"lat"`
	Long       string    `json:"long"`
	Width      float64   `json:"width"`
	LocationID uuid.UUID `json:"location_id"`
}

type LocationProperties []LocationProperty
