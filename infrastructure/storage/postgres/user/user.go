package user

import (
	"database/sql"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

var (
	_psqlGetById = `select * from domain.users where id = $1`
	_psqlGetAll  = `select * from domain.users`
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) GetStorageById(id uuid.UUID) (*model.User, error) {
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

func (u User) GetStorageAll() (model.Users, error) {
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
