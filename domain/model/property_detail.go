package model

import "github.com/google/uuid"

type PropertyDetail struct {
	ID          uuid.UUID `json:"id"`
	Rooms       int       `json:"rooms"`
	Bathrooms   int       `json:"bathrooms"`
	Yard        int       `json:"yard"`
	Terrace     bool      `json:"terrace"`
	LivingRoom  int       `json:"living_room"`
	LaundryRoom int       `json:"laundry_room"`
	Kitchen     int       `json:"kitchen"`
	Garage      bool      `json:"garage"`
}

type PropertiesDetail []PropertyDetail
