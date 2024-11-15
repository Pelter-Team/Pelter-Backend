package transaction

import (
	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/user"
	"context"
)

type (
	transactionUsecase struct {
		transactionRepo TransactionRepository
		userRepo        user.UserRepository
	}

	TransactionUsecase interface {
		Buy(ctx context.Context, txn *entity.Transaction) error
		FindById(ctx context.Context, id uint) (*entity.Transaction, error)
		ListAll(ctx context.Context) ([]*entity.Transaction, error)
	}
)

func NewTransactionUsecase(transactionRepo TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
	}
}

func (u *transactionUsecase) Buy(ctx context.Context, txn *entity.Transaction) error {
	return u.transactionRepo.CreateTransaction(ctx, txn)
}

func (u *transactionUsecase) FindById(ctx context.Context, id uint) (*entity.Transaction, error) {
	return u.transactionRepo.FindById(ctx, id)
}

func (u *transactionUsecase) ListAll(ctx context.Context) ([]*entity.Transaction, error) {
	return u.transactionRepo.ListAll(ctx)
}
