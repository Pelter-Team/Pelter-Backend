package product

import (
	"context"
)

type (
	productUsecase struct {
		productRepo ProductRepository
	}
	ProductUsecase interface {
		InsertProduct(pctx context.Context) error
	}
)

func NewProductUsecase(productRepo ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) InsertProduct(pctx context.Context) error {
	_ = u.productRepo.InsertProduct(pctx)
	return nil
}
