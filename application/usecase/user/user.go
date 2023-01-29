package user

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/user"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type User struct {
	storage user.Storage
}

func New(storage user.Storage) User {
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
