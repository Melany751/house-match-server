package user

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

var (
	_psqlGetById = `select * from domain.users where id = $1 `
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) GetById(id uuid.UUID) (*model.User, error) {
	query := _psqlGetById

	fmt.Println("################")
	fmt.Println("query: ", query)
	var m model.User

	err := u.db.QueryRow(query, id).Scan(&m.ID, &m.User, &m.Password, &m.Email, &m.Theme)
	if err != nil {
		panic(err)
	}
	//defer stmt.Close()
	//
	//m, err := u.scanRow(stmt.QueryRow())
	//if err != nil {
	//	return nil, err
	//}

	fmt.Println("model: ", m)
	fmt.Println("model: ", &m)

	return &m, nil
}

func (u User) scanRow(s pgx.Row) (model.User, error) {
	m := model.User{}

	themeNull := sql.NullString{}

	err := s.Scan(
		&m.ID,
		&m.User,
		&m.Password,
		&m.Email,
		themeNull,
	)
	if err != nil {
		return model.User{}, err
	}

	m.Theme = themeNull.String

	return m, nil
}
