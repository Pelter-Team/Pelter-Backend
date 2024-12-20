package transaction

import (
	"Pelter_backend/internal/dto"
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
		CreateTransaction(ctx context.Context, txn *entity.Transaction) (entity.Transaction, error)
		GetTransactions(ctx context.Context) ([]*dto.TransactionWithProductResponse, error)
		GetTransactionByID(ctx context.Context, id uint) (*entity.Transaction, error)
		GetTransactionsByUserID(ctx context.Context, id uint) ([]*entity.Transaction, error)
	}
)

func NewTransactionUsecase(transactionRepo TransactionRepository, userRepo user.UserRepository, productRepo product.ProductRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		productRepo:     productRepo,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, txn *entity.Transaction) (entity.Transaction, error) {
	product, err := u.productRepo.GetProductByID(ctx, txn.ProductID)
	if err != nil {
		return entity.Transaction{}, errors.New("Product not found")
	}
	if product.IsSold {
		return entity.Transaction{}, errors.New("Product already sold")
	}
	if product.UserID == txn.BuyerID {
		return entity.Transaction{}, errors.New("You cannot buy your own product")
	}
	txn.SellerID = product.UserID
	txn.Amount = product.Price
	if err := u.transactionRepo.CreateTransaction(ctx, txn); err != nil {
		return entity.Transaction{}, err
	}
	product.IsSold = true
	if err := u.productRepo.UpdateProduct(ctx, &product, uint(product.ID), uint(product.UserID)); err != nil {
		return entity.Transaction{}, err
	}
	return *txn, nil
}

func (u *transactionUsecase) GetTransactions(ctx context.Context) ([]*dto.TransactionWithProductResponse, error) {
	_transactions, err := u.transactionRepo.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}
	transactions := make([]*dto.TransactionWithProductResponse, 0, len(_transactions))
	for _, txn := range _transactions {
		transactions = append(transactions, &dto.TransactionWithProductResponse{
			ID:         txn.ID,
			ProductID:  txn.ProductID,
			BuyerID:    txn.BuyerID,
			SellerID:   txn.SellerID,
			Amount:     uint(txn.Amount),
			CreatedAt:  txn.CreatedAt.String(),
			Price:      txn.Product.Price,
			IsVerified: txn.Product.IsVerified,
			IsSold:     txn.Product.IsSold,
		})
	}
	return transactions, nil
}

func (u *transactionUsecase) GetTransactionByID(ctx context.Context, id uint) (*entity.Transaction, error) {
	transaction, err := u.transactionRepo.FindByTransactionID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (u *transactionUsecase) GetTransactionsByUserID(ctx context.Context, id uint) ([]*entity.Transaction, error) {
	transactions, err := u.transactionRepo.FindByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, errors.New("No transactions found")
	}
	return transactions, nil
}
