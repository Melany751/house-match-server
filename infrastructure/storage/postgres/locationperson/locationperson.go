package locationperson

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.location_persons"
)

var (
	_psqlGetById = `SELECT * FROM domain.location_persons WHERE id = $1`
	_psqlGetAll  = `SELECT * FROM domain.location_persons`
	_psqlInsert  = `INSERT INTO domain.location_persons (id, "country", "city", "province", "district") VALUES ($1, $2, $3, $4, $5)`
	_psqlUpdate  = `UPDATE domain.location_persons SET "country"=$2, "city"=$3, "province"=$4, "district"=$5 WHERE id=$1`
	_psqlDelete  = `DELETE FROM domain.location_persons WHERE id=$1`
)

type LocationPerson struct {
	db *sql.DB
}

func New(db *sql.DB) LocationPerson {
	return LocationPerson{db}
}

func (u LocationPerson) GetByIdStorage(id uuid.UUID) (*model.Location, error) {
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

func (u LocationPerson) GetAllStorage() (model.Locations, error) {
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

	var ms model.Locations
	var m model.Location

	for rows.Next() {
		m, err = u.scanRow(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u LocationPerson) CreateStorage(locationPerson model.Location) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	locationPerson.ID = newId

	args := u.readModelLocationPerson(locationPerson)

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

func (u LocationPerson) UpdateStorage(id uuid.UUID, locationPerson model.Location) (bool, error) {
	locationPerson.ID = id

	args := u.readModelLocationPerson(locationPerson)

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

func (u LocationPerson) DeleteStorage(id uuid.UUID) (bool, error) {
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

func (u LocationPerson) readModelLocationPerson(locationPerson model.Location) []any {
	v := reflect.ValueOf(locationPerson)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (u LocationPerson) scanRow(s pgx.Row) (model.Location, error) {
	m := model.Location{}

	err := s.Scan(
		&m.ID,
		&m.Country,
		&m.City,
		&m.Province,
		&m.District,
	)
	if err != nil {
		return model.Location{}, err
	}

	return m, nil
}
