package location

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseLocation interface {
	GetById(id uuid.UUID) (*model.Location, error)
	GetAll() (model.Locations, error)
	Create(m model.Location) (*model.CreateOutput, error)
	Update(id uuid.UUID, user model.Location) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
