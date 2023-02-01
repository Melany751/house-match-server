package role

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseRole interface {
	GetById(id uuid.UUID) (*model.Role, error)
	GetAll() (model.Roles, error)
	Create(m model.Role) (*uuid.UUID, error)
	Update(id uuid.UUID, user model.Role) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}
