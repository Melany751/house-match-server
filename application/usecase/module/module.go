package module

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/module"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Module struct {
	storage module.Storage
}

func New(storage module.Storage) Module {
	return Module{storage}
}

func (r Module) GetById(id uuid.UUID) (*model.Module, error) {
	role, err := r.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("role.storage.GetById(): %w", err)
	}

	return role, nil
}

func (r Module) GetAll() (model.Modules, error) {
	roles, err := r.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("role.storage.GetAll(): %w", err)
	}

	return roles, nil
}

func (r Module) Create(role model.Module) (*uuid.UUID, error) {
	id, err := r.storage.CreateStorage(role)
	if err != nil {
		return nil, fmt.Errorf("role.storage.Create(): %w", err)
	}

	return id, nil
}

func (r Module) Update(id uuid.UUID, role model.Module) (bool, error) {
	created, err := r.storage.UpdateStorage(id, role)
	if err != nil {
		return false, fmt.Errorf("role.storage.Update(): %w", err)
	}

	return created, nil
}

func (r Module) Delete(id uuid.UUID) (bool, error) {
	deleted, err := r.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("role.storage.Delete(): %w", err)
	}

	return deleted, nil
}
