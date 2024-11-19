package transaction

import (
	"Pelter_backend/internal/entity"
	"fmt"

	"context"

	"gorm.io/gorm"
)

type (
	transactionRepository struct {
		Db *gorm.DB
	}

	TransactionRepository interface {
		CreateTransaction(ctx context.Context, txn *entity.Transaction) error
		GetTransactions(ctx context.Context) ([]*entity.Transaction, error)
		FindByUserID(ctx context.Context, userID uint) ([]*entity.Transaction, error)
		FindByTransactionID(ctx context.Context, ID uint) (*entity.Transaction, error)
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

func (r *transactionRepository) GetTransactions(pctx context.Context) ([]*entity.Transaction, error) {
	var txns []*entity.Transaction
	if err := r.transactionTable(pctx).Find(&txns).Error; err != nil {
		return nil, err
	}
	return txns, nil
}

func (r *transactionRepository) FindByTransactionID(pctx context.Context, ID uint) (*entity.Transaction, error) {
	var txn entity.Transaction
	if err := r.transactionTable(pctx).Where("id = ?", ID).First(&txn).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("transaction with ID %d not found", ID)
		}
		return nil, err
	}
	return &txn, nil
}

func (r *transactionRepository) FindByUserID(ctx context.Context, userID uint) ([]*entity.Transaction, error) {
	var txns []*entity.Transaction
	if err := r.transactionTable(ctx).Where("buyer_id = ? OR seller_id = ?", userID, userID).Find(&txns).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("transaction with ID %d not found for both buyer and seller", userID)
		}
		return nil, err
	}
	return txns, nil
}
