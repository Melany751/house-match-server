package property

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageProperty interface {
	GetStorageById(id uuid.UUID) (*model.Property, error)
	GetStorageAll() (model.Properties, error)
	CreateStorage(m model.Property) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, model model.Property) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}