package module

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/module"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Module struct {
	storage module.StorageModule
}

func New(storage module.StorageModule) Module {
	return Module{storage}
}

func (m Module) GetById(id uuid.UUID) (*model.Module, error) {
	role, err := m.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("role.storage.GetById(): %w", err)
	}

	return role, nil
}

func (m Module) GetAll() (model.Modules, error) {
	roles, err := m.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("role.storage.GetAll(): %w", err)
	}

	return roles, nil
}

func (m Module) Create(role model.Module) (*uuid.UUID, error) {
	id, err := m.storage.CreateStorage(role)
	if err != nil {
		return nil, fmt.Errorf("role.storage.Create(): %w", err)
	}

	return id, nil
}

func (m Module) Update(id uuid.UUID, role model.Module) (bool, error) {
	created, err := m.storage.UpdateStorage(id, role)
	if err != nil {
		return false, fmt.Errorf("role.storage.Update(): %w", err)
	}

	return created, nil
}

func (m Module) Delete(id uuid.UUID) (bool, error) {
	deleted, err := m.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("role.storage.Delete(): %w", err)
	}

	return deleted, nil
}
