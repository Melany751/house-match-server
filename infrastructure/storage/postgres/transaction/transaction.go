package transaction

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.transactions"
)

var (
	_psqlGetById = `SELECT
    				t.id,
    				t.cost,
    				t.currency,
    				t.date_vip,
    				t.date_post,
    				t.date_update,
    				t.available,
    				t.type,
    				t.date_start,
    				t.date_end,
    				p.id,
    				p.user_id,
    				p.location_id,
    				p.description,
    				p.type,
    				p.length,
    				p.width,
    				p.area,
    				p.floor,
    				p.number_of_floors,
    				p.rooms,
					p.bathrooms,
					p.yard,
					p.terrace,
					p.living_room,
					p.laundry_room,
					p.kitchen,
					p.garage 
					FROM domain.transactions t
         			INNER JOIN domain.properties p ON t.property_id=p.id
         			WHERE t.id = $1`
	_psqlGetAll = `SELECT 
    				t.id,
    				t.cost,
    				t.currency,
    				t.date_vip,
    				t.date_post,
    				t.date_update,
    				t.available,
    				t.type,
    				t.date_start,
    				t.date_end,
    				p.id,
    				p.user_id,
    				p.location_id,
    				p.description,
    				p.type,
    				p.length,
    				p.width,
    				p.area,
    				p.floor,
    				p.number_of_floors,
    				p.rooms,
					p.bathrooms,
					p.yard,
					p.terrace,
					p.living_room,
					p.laundry_room,
					p.kitchen,
					p.garage 
					FROM domain.transactions t
         			INNER JOIN domain.properties p ON t.property_id=p.id`
	_psqlInsert = `INSERT INTO domain.transactions (id, property_id,cost,currency,date_vip,date_post,date_update,available,"type",date_start,date_end) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_psqlUpdate = `UPDATE domain.transactions SET property_id=$2,cost=$3,currency=$4,date_vip=$5,date_post=$6,date_update=$7,available=$8,"type"=$9,date_start=$10,date_end=$11 WHERE id=$1`
	_psqlDelete = `DELETE FROM domain.transactions WHERE id=$1`
)

type Transaction struct {
	db *sql.DB
}

func New(db *sql.DB) Transaction {
	return Transaction{db}
}

func (r Transaction) GetStorageById(id uuid.UUID) (*model.TransactionSecondLevel, error) {
	args := []any{id}

	stmt, err := r.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := r.scanRowWithProperty(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r Transaction) GetStorageAll() (model.TransactionsSecondLevel, error) {
	stmt, err := r.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.TransactionsSecondLevel
	var m model.TransactionSecondLevel

	for rows.Next() {
		m, err = r.scanRowWithProperty(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (r Transaction) CreateStorage(transaction model.Transaction) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	transaction.ID = newId

	args := r.readModelTransaction(transaction)

	stmt, err := r.db.Prepare(_psqlInsert)
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

func (r Transaction) UpdateStorage(id uuid.UUID, transaction model.Transaction) (bool, error) {
	transaction.ID = id

	args := r.readModelTransaction(transaction)

	stmt, err := r.db.Prepare(_psqlUpdate)
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

func (r Transaction) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := r.db.Prepare(_psqlDelete)
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

func (r Transaction) readModelTransaction(transaction model.Transaction) []any {
	v := reflect.ValueOf(transaction)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (r Transaction) scanRow(s pgx.Row) (model.Transaction, error) {
	m := model.Transaction{}

	err := s.Scan(
		&m.ID,
		&m.PropertyID,
		&m.Cost,
		&m.Currency,
		&m.DateVIP,
		&m.DatePost,
		&m.DateUpdate,
		&m.Available,
		&m.Type,
		&m.DateStart,
		&m.DateEnd,
	)
	if err != nil {
		return model.Transaction{}, err
	}

	return m, nil
}

func (r Transaction) scanRowWithProperty(s pgx.Row) (model.TransactionSecondLevel, error) {
	m := model.TransactionSecondLevel{}

	err := s.Scan(
		&m.ID,
		&m.Cost,
		&m.Currency,
		&m.DateVIP,
		&m.DatePost,
		&m.DateUpdate,
		&m.Available,
		&m.Type,
		&m.DateStart,
		&m.DateEnd,
		&m.Property.ID,
		&m.Property.UserID,
		&m.Property.LocationID,
		&m.Property.Description,
		&m.Property.Type,
		&m.Property.Length,
		&m.Property.Width,
		&m.Property.Area,
		&m.Property.Floor,
		&m.Property.NumberOfFloors,
		&m.Property.Rooms,
		&m.Property.Bathrooms,
		&m.Property.Yard,
		&m.Property.Terrace,
		&m.Property.LivingRoom,
		&m.Property.LaundryRoom,
		&m.Property.Kitchen,
		&m.Property.Garage,
	)
	if err != nil {
		return model.TransactionSecondLevel{}, err
	}

	return m, nil
}
