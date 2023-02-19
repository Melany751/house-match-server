package locationperson

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseLocationPerson interface {
	GetById(id uuid.UUID) (*model.LocationPerson, error)
	GetAll() (model.LocationPersons, error)
	Create(m model.LocationPerson) (*model.CreateOutput, error)
	Update(id uuid.UUID, user model.LocationPerson) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
