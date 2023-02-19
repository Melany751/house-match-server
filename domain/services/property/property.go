package property

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseModule interface {
	GetById(id uuid.UUID) (*model.PropertyOutput, error)
	GetAll() (model.PropertiesOutput, error)
	Create(m model.Property) (*model.CreateOutput, error)
	Update(id uuid.UUID, model model.Property) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (bool, error)
}
