package role

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/role"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Role struct {
	storage role.Storage
}

func New(storage role.Storage) Role {
	return Role{storage}
}

func (r Role) GetById(id uuid.UUID) (*model.Role, error) {
	role, err := r.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("role.storage.GetById(): %w", err)
	}

	return role, nil
}

func (r Role) GetAll() (model.Roles, error) {
	roles, err := r.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("role.storage.GetAll(): %w", err)
	}

	return roles, nil
}

func (r Role) Create(role model.Role) (*uuid.UUID, error) {
	id, err := r.storage.CreateStorage(role)
	if err != nil {
		return nil, fmt.Errorf("role.storage.Create(): %w", err)
	}

	return id, nil
}

func (r Role) Update(id uuid.UUID, role model.Role) (bool, error) {
	created, err := r.storage.UpdateStorage(id, role)
	if err != nil {
		return false, fmt.Errorf("role.storage.Update(): %w", err)
	}

	return created, nil
}

func (r Role) Delete(id uuid.UUID) (bool, error) {
	deleted, err := r.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("role.storage.Delete(): %w", err)
	}

	return deleted, nil
}
