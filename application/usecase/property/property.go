package property

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/property"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Property struct {
	storage property.StorageProperty
}

func New(storage property.StorageProperty) Property {
	return Property{storage}
}

func (p Property) GetById(id uuid.UUID) (*model.Property, error) {
	property, err := p.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("property.storage.GetById(): %w", err)
	}

	return property, nil
}

func (p Property) GetAll() (model.Properties, error) {
	properties, err := p.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("property.storage.GetAll(): %w", err)
	}

	return properties, nil
}

func (p Property) Create(property model.Property) (*uuid.UUID, error) {
	id, err := p.storage.CreateStorage(property)
	if err != nil {
		return nil, fmt.Errorf("property.storage.Create(): %w", err)
	}

	return id, nil
}

func (p Property) Update(id uuid.UUID, property model.Property) (bool, error) {
	created, err := p.storage.UpdateStorage(id, property)
	if err != nil {
		return false, fmt.Errorf("property.storage.Update(): %w", err)
	}

	return created, nil
}

func (p Property) Delete(id uuid.UUID) (bool, error) {
	deleted, err := p.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("property.storage.Delete(): %w", err)
	}

	return deleted, nil
}
