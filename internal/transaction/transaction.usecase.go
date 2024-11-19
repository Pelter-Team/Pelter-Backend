package transaction

import (
	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/product"
	"Pelter_backend/internal/user"
	"context"
	"errors"
)

type (
	transactionUsecase struct {
		transactionRepo TransactionRepository
		userRepo        user.UserRepository
		productRepo     product.ProductRepository
	}

	TransactionUsecase interface {
		Buy(ctx context.Context, txn *entity.Transaction) error
		FindById(ctx context.Context, id uint) (*entity.Transaction, error)
		ListAll(ctx context.Context) ([]*entity.Transaction, error)
		ListAllByUserId(ctx context.Context, id uint) ([]*entity.Transaction, error)
	}
)

func NewTransactionUsecase(transactionRepo TransactionRepository, userRepo UserRepository, productRepo ProductRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		// productRepo:     productRepo,
	}
}

func (u *transactionUsecase) Buy(ctx context.Context, txn *entity.Transaction) error {
	// Ensure the product exists
	if _, err := u.productRepo.FindByID(ctx, txn.ProductID); err != nil {
		return errors.New("product not found")
	}
	// Ensure the buyer exists
	if seller_id, err := u.userRepo.FindByID(ctx, txn.BuyerID); err != nil {
		return errors.New("buyer not found")
	}

	// Ensure the seller exists
	if _, err := u.userRepo.FindByID(ctx, txn.SellerID); err != nil {
		return errors.New("seller not found")
	}

	return u.transactionRepo.CreateTransaction(ctx, txn)
}

func (u *transactionUsecase) FindById(ctx context.Context, id uint) (*entity.Transaction, error) {
	return u.transactionRepo.FindById(ctx, id)
}

func (u *transactionUsecase) ListAll(ctx context.Context) ([]*entity.Transaction, error) {
	return u.transactionRepo.ListAll(ctx)
}

func (u *transactionUsecase) ListAllByUserId(ctx context.Context, id uint) ([]*entity.Transaction, error) {
	return u.transactionRepo.ListAllByUserId(ctx, id)
}
