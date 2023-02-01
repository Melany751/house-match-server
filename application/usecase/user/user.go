package user

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/user"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type User struct {
	storage user.StorageUser
}

func New(storage user.StorageUser) User {
	return User{storage}
}

func (u User) GetById(id uuid.UUID) (*model.User, error) {
	user, err := u.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetById(): %w", err)
	}

	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetAll(): %w", err)
	}

	return users, nil
}

func (u User) Create(user model.User) (*uuid.UUID, error) {
	id, err := u.storage.CreateStorage(user)
	if err != nil {
		return nil, fmt.Errorf("user.storage.Create(): %w", err)
	}

	return id, nil
}

func (u User) Update(id uuid.UUID, user model.User) (bool, error) {
	created, err := u.storage.UpdateStorage(id, user)
	if err != nil {
		return false, fmt.Errorf("user.storage.Update(): %w", err)
	}

	return created, nil
}

func (u User) Delete(id uuid.UUID) (bool, error) {
	deleted, err := u.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("user.storage.Delete(): %w", err)
	}

	return deleted, nil
}
