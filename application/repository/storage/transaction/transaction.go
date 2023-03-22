package transaction

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageTransaction interface {
	GetStorageById(id uuid.UUID) (*model.TransactionSecondLevel, error)
	GetStorageAll() (model.TransactionsSecondLevel, error)
	CreateStorage(m model.Transaction) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, role model.Transaction) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
