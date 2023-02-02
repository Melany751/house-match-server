package user

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.users"
)

var (
	_psqlGetById = `SELECT * FROM domain.users WHERE id = $1`
	_psqlGetAll  = `SELECT * FROM domain.users`
	_psqlInsert  = `INSERT INTO domain.users (id, "user", "password", "email", "theme") VALUES ($1, $2, $3, $4, $5)`
	_psqlUpdate  = `UPDATE domain.users SET "user"=$2, "password"=$3, "email"=$4, "theme"=$5 WHERE id=$1`
	_psqlDelete  = `DELETE FROM domain.users WHERE id=$1`
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) GetByIdStorage(id uuid.UUID) (*model.User, error) {
	args := []any{id}

	stmt, err := u.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := u.scanRow(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (u User) GetAllStorage() (model.Users, error) {
	stmt, err := u.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Users
	var m model.User

	for rows.Next() {
		m, err = u.scanRow(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) CreateStorage(user model.User) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	user.ID = newId

	args := u.readModelUser(user)

	stmt, err := u.db.Prepare(_psqlInsert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return &newId, nil
}

func (u User) UpdateStorage(id uuid.UUID, user model.User) (bool, error) {
	user.ID = id

	args := u.readModelUser(user)

	stmt, err := u.db.Prepare(_psqlUpdate)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("nothing rows updated, error: %v", err)
	}

	return true, nil
}

func (u User) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := u.db.Prepare(_psqlDelete)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("nothing rows delete, error: %v", err)
	}

	return true, nil
}

func (u User) readModelUser(user model.User) []any {
	v := reflect.ValueOf(user)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (u User) scanRow(s pgx.Row) (model.User, error) {
	m := model.User{}

	err := s.Scan(
		&m.ID,
		&m.User,
		&m.Password,
		&m.Email,
		&m.Theme,
	)
	if err != nil {
		return model.User{}, err
	}

	return m, nil
}
