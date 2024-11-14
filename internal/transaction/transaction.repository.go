package transaction

import (
	"Pelter_backend/internal/entity"

	"context"

	"gorm.io/gorm"
)

type (
	transactionRepository struct {
		Db *gorm.DB
	}

	TransactionRepository interface {
		CreateTransaction(ctx context.Context, txn *entity.Transaction) error
		FindById(ctx context.Context, id uint) (*entity.Transaction, error)
		ListAll(ctx context.Context) ([]*entity.Transaction, error)
		ListAllByUserId(ctx context.Context, id uint) ([]*entity.Transaction, error)
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
	if err := r.transactionTable(pctx).Where("product_id = ?", txn.ProductID).First(&txn).Error; err == nil {
		return gorm.ErrDuplicatedKey
	}
	return r.transactionTable(pctx).Create(txn).Error
}

func (r *transactionRepository) FindById(pctx context.Context, id uint) (*entity.Transaction, error) {
	var txn entity.Transaction
	if err := r.transactionTable(pctx).Where("id = ?", id).First(&txn).Error; err != nil {
		return nil, err
	}
	return &txn, nil
}

func (r *transactionRepository) ListAll(pctx context.Context) ([]*entity.Transaction, error) {
	var txns []*entity.Transaction
	if err := r.transactionTable(pctx).Find(&txns).Error; err != nil {
		return nil, err
	}
	return txns, nil
}

func (r *transactionRepository) ListAllByUserId(pctx context.Context, id uint) ([]*entity.Transaction, error) {
	var txns []*entity.Transaction
	if err := r.transactionTable(pctx).Where("buyer_id = ?", id).Find(&txns).Error; err != nil {
		return nil, err
	}
	return txns, nil
}
