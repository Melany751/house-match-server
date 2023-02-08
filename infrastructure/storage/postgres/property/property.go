package property

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.properties"
)

var (
	_psqlGetById = `SELECT * FROM domain.properties WHERE id = $1`
	_psqlGetAll  = `SELECT * FROM domain.properties`
	_psqlInsert  = `INSERT INTO domain.properties (id, "name", "description", "icon", "order") VALUES ($1, $2, $3, $4, $5)`
	_psqlUpdate  = `UPDATE domain.properties SET "name"=$2, "description"=$3, "icon"=$4, "order"=$5 WHERE id=$1`
	_psqlDelete  = `DELETE FROM domain.properties WHERE id=$1`
)

type Property struct {
	db *sql.DB
}

func New(db *sql.DB) Property {
	return Property{db}
}

func (p Property) GetStorageById(id uuid.UUID) (*model.Property, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := p.scanRow(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (p Property) GetStorageAll() (model.Properties, error) {
	stmt, err := p.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Properties
	var m model.Property

	for rows.Next() {
		m, err = p.scanRow(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Property) CreateStorage(property model.Property) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	property.ID = newId

	args := p.readModelProperty(property)

	stmt, err := p.db.Prepare(_psqlInsert)
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

func (p Property) UpdateStorage(id uuid.UUID, property model.Property) (bool, error) {
	property.ID = id

	args := p.readModelProperty(property)

	stmt, err := p.db.Prepare(_psqlUpdate)
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

func (p Property) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlDelete)
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

func (p Property) readModelProperty(property model.Property) []any {
	v := reflect.ValueOf(property)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (p Property) scanRow(s pgx.Row) (model.Property, error) {
	m := model.Property{}

	err := s.Scan(
		&m.ID,
		&m.UserID,
		&m.Description,
		&m.Type,
		&m.Length,
		&m.Width,
		&m.Area,
		&m.Floor,
		&m.NumberOfFloor,
	)
	if err != nil {
		return model.Property{}, err
	}

	return m, nil
}
