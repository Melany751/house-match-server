package user

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseUser interface {
	GetById(id uuid.UUID) (*model.User, error)
	GetAll() (model.Users, error)
	Create(m model.User) (*uuid.UUID, error)
	Update(id uuid.UUID, user model.User) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}
