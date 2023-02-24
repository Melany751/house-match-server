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
	_psqlGetById = `SELECT 
    				p.id,
    				p.description,
    				p.type,
    				p.length,
    				p.width,
    				p.area,
    				p.floor,
    				p.number_of_floors,
    				u.id, 
    				u.user, 
    				u.password, 
    				u.email, 
    				u.theme
					FROM domain.properties p INNER JOIN domain.users u ON p.user_id=u.id
					WHERE p.id = $1`
	_psqlGetAll = `SELECT 
    				p.id,
    				p.description,
    				p.type,
    				p.length,
    				p.width,
    				p.area,
    				p.floor,
    				p.number_of_floors,
    				u.id, 
    				u.user, 
    				u.password, 
    				u.email, 
    				u.theme FROM domain.properties p INNER JOIN domain.users u ON p.user_id=u.id`
	_psqlInsert = `INSERT INTO domain.properties (id, "user_id", "description", "type", "length","width","area","floor","number_of_floors") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_psqlUpdate = `UPDATE domain.properties SET "user_id"=$2, "description"=$3, "type"=$4, "length"=$5, "width"=$6, "area"=$7, "floor"=$8, "number_of_floors"=$9 WHERE id=$1`
	_psqlDelete = `DELETE FROM domain.properties WHERE id=$1`
)

type Property struct {
	db *sql.DB
}

func New(db *sql.DB) Property {
	return Property{db}
}

func (p Property) GetStorageById(id uuid.UUID) (*model.PropertySecondLevel, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := p.scanRowWithUser(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (p Property) GetStorageAll() (model.PropertiesSecondLevel, error) {
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

	var ms model.PropertiesSecondLevel
	var m model.PropertySecondLevel

	for rows.Next() {
		m, err = p.scanRowWithUser(rows)
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

func (p Property) scanRowWithUser(s pgx.Row) (model.PropertySecondLevel, error) {
	m := model.PropertySecondLevel{}
	err := s.Scan(
		&m.ID,
		&m.Description,
		&m.Type,
		&m.Length,
		&m.Width,
		&m.Area,
		&m.Floor,
		&m.NumberOfFloor,
		&m.User.ID,
		&m.User.User,
		&m.User.Password,
		&m.User.Email,
		&m.User.Theme,
	)
	if err != nil {
		return model.PropertySecondLevel{}, err
	}

	return m, nil
}
