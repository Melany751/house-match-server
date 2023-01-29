package user

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Storage interface {
	GetStorageById(id uuid.UUID) (*model.User, error)
	GetStorageAll() (model.Users, error)
	//Create(m model.User) (bool, error)
	//Update(id uuid.UUID, user model.User) (model.Users, error)
	//Delete(id uuid.UUID) (bool, error)
}