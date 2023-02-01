package view

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/view"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type View struct {
	storage view.StorageView
}

func New(storage view.StorageView) View {
	return View{storage}
}

func (v View) GetById(id uuid.UUID) (*model.ViewOutput, error) {
	view, err := v.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("view.storage.GetById(): %w", err)
	}
	return view, nil
}

func (v View) GetAll() (model.ViewsOutput, error) {
	views, err := v.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("view.storage.GetAll(): %w", err)
	}
	return views, nil
}

func (v View) Create(view model.View) (*uuid.UUID, error) {
	id, err := v.storage.CreateStorage(view)
	if err != nil {
		return nil, fmt.Errorf("view.storage.Create(): %w", err)
	}

	return id, nil
}

func (v View) Update(id uuid.UUID, view model.View) (bool, error) {
	created, err := v.storage.UpdateStorage(id, view)
	if err != nil {
		return false, fmt.Errorf("view.storage.Update(): %w", err)
	}

	return created, nil
}

func (v View) Delete(id uuid.UUID) (bool, error) {
	deleted, err := v.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("view.storage.Delete(): %w", err)
	}

	return deleted, nil
}
