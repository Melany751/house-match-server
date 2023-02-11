package module

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseModule interface {
	GetById(id uuid.UUID) (*model.Module, error)
	GetAll() (model.Modules, error)
	Create(m model.Module) (*uuid.UUID, error)
	Update(id uuid.UUID, model model.Module) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}