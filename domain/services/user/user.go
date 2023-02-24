package user

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseUser interface {
	GetById(id uuid.UUID) (*model.UserOutput, error)
	GetAll() (model.UsersOutput, error)
	Create(m model.User) (*model.CreateOutput, error)
	Update(id uuid.UUID, user model.User) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
