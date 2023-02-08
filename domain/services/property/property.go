package property

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseModule interface {
	GetById(id uuid.UUID) (*model.Property, error)
	GetAll() (model.Properties, error)
	Create(m model.Property) (*uuid.UUID, error)
	Update(id uuid.UUID, model model.Property) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}
