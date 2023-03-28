package transaction

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseTransaction interface {
	GetById(id uuid.UUID) (*model.TransactionSecondLevel, error)
	GetAll() (model.TransactionsSecondLevel, error)
	Create(m model.Transaction) (*model.CreateOutput, error)
	Update(id uuid.UUID, model model.Transaction) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (bool, error)
}
