package transaction

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/transaction"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	storage transaction.StorageTransaction
}

func New(storage transaction.StorageTransaction) Transaction {
	return Transaction{storage}
}

func (r Transaction) GetById(id uuid.UUID) (*model.TransactionSecondLevel, error) {
	transaction, err := r.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("transaction.storage.GetById(): %w", err)
	}

	return transaction, nil
}

func (r Transaction) GetAll() (model.TransactionsSecondLevel, error) {
	transactions, err := r.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("transaction.storage.GetAll(): %w", err)
	}

	return transactions, nil
}

func (r Transaction) Create(transaction model.Transaction) (*model.CreateOutput, error) {
	transaction.DateVIP = time.Now().Add((time.Hour * 24) * -1)
	transaction.DatePost = time.Now()
	transaction.DateUpdate = time.Now()
	transaction.Available = true
	id, err := r.storage.CreateStorage(transaction)
	if err != nil {
		return nil, fmt.Errorf("transaction.storage.Create(): %w", err)
	}
	var m model.CreateOutput
	m.Id = id
	return &m, nil
}

func (r Transaction) Update(id uuid.UUID, transaction model.Transaction) (*model.UpdateOutput, error) {
	updated, err := r.storage.UpdateStorage(id, transaction)
	if err != nil {
		return nil, fmt.Errorf("transaction.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = updated
	return &m, nil
}

func (r Transaction) Delete(id uuid.UUID) (bool, error) {
	deleted, err := r.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("transaction.storage.Delete(): %w", err)
	}

	return deleted, nil
}
