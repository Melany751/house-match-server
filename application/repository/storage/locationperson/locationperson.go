package locationperson

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageLocationPerson interface {
	GetByIdStorage(id uuid.UUID) (*model.LocationPerson, error)
	GetAllStorage() (model.LocationPersons, error)
	CreateStorage(m model.LocationPerson) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, user model.LocationPerson) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
