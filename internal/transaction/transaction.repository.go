package transaction

import (
	"context"

	"gorm.io/gorm"

	"Pelter_backend/internal/entity"
)

type (
	transactionRepository struct {
		Db *gorm.DB
	}

	TransactionRepository interface {
		CreateTransaction(ctx context.Context, txn *entity.Transaction) error
		FindById(ctx context.Context, id uint) (*entity.Transaction, error)
		ListAll(ctx context.Context) error
	}
)

func (r *transactionRepository) transactionTable(pctx context.Context) *gorm.DB {
	return r.Db.Table("transactions").WithContext(pctx)
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		Db: db,
	}
}

func (r *transactionRepository) CreateTransaction(pctx context.Context, txn *entity.Transaction) error {
	return r.transactionTable(pctx).Create(txn).Error
}

func (r *transactionRepository) FindById(pctx context.Context, id uint) (*entity.Transaction, error) {
	var txn entity.Transaction
	if err := r.transactionTable(pctx).Where("id = ?", id).First(&txn).Error; err != nil {
		return nil, err
	}
	return &txn, nil
}

func (r *transactionRepository) ListAll(pctx context.Context) error {
	return r.transactionTable(pctx).Find(&[]entity.Transaction{}).Error
}
