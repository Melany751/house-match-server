package user

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Storage interface {
	GetStorageById(id uuid.UUID) (*model.User, error)
	GetStorageAll() (model.Users, error)
	CreateStorage(m model.User) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, user model.User) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
