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

func (u User) GetById(id uuid.UUID) (*model.UserOutput, error) {
	user, err := u.storage.GetByIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetById(): %w", err)
	}

	return user, nil
}

func (u User) GetAll() (model.UsersOutput, error) {
	users, err := u.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetAll(): %w", err)
	}

	return users, nil
}

func (u User) GetAllWithRoles() (model.UsersWithRolesOutput, error) {
	users, err := u.storage.GetAllWithRolesStorage()
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetAllWithRoles(): %w", err)
	}

	return users, nil
}

func (u User) Create(user model.User) (*model.CreateOutput, error) {
	id, err := u.storage.CreateStorage(user)
	if err != nil {
		return nil, fmt.Errorf("user.use.Create(): %w", err)
		//return nil, fmt.Errorf("user.use.Create(): %s", err.Error())
	}

	var m model.CreateOutput
	m.Id = id

	return &m, nil
}

func (u User) Update(id uuid.UUID, user model.User) (*model.UpdateOutput, error) {
	created, err := u.storage.UpdateStorage(id, user)
	if err != nil {
		return nil, fmt.Errorf("user.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = created

	return &m, nil
}

func (u User) Delete(id uuid.UUID) (*model.DeleteOutput, error) {
	deleted, err := u.storage.DeleteStorage(id)
	if err != nil {
		return nil, fmt.Errorf("user.storage.Delete(): %w", err)
	}

	var m model.DeleteOutput
	m.Deleted = deleted

	return &m, nil
}
