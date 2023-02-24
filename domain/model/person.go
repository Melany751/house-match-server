package model

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	ID               uuid.UUID `json:"id"`
	DocumentType     string    `json:"document_type"`
	Document         string    `json:"document"`
	Names            string    `json:"names"`
	Lastname         string    `json:"lastname"`
	MLastname        string    `json:"m_lastname"`
	Phone            string    `json:"phone"`
	Gender           string    `json:"gender"`
	MaritalStatus    string    `json:"marital_status"`
	DateBirth        time.Time `json:"date_birth"`
	LocationPersonID uuid.UUID `json:"location_person_id"`
}

type Persons []Person

type PersonOutput struct {
	ID             uuid.UUID `json:"id"`
	DocumentType   string    `json:"document_type"`
	Document       string    `json:"document"`
	Names          string    `json:"names"`
	Lastname       string    `json:"lastname"`
	MLastname      string    `json:"m_lastname"`
	Phone          string    `json:"phone"`
	Gender         string    `json:"gender"`
	MaritalStatus  string    `json:"marital_status"`
	DateBirth      time.Time `json:"date_birth"`
	LocationPerson Location  `json:"location_person"`
}

type PersonsOutput []PersonOutput
