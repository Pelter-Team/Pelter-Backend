package product

import (
	"context"
)

type (
	usecase struct {
		productRepo ProductRepository
	}
	ProductUsecase interface {
		InsertProduct(pctx context.Context) error
	}
)

func NewProductUsecase(productRepo ProductRepository) ProductUsecase {
	return &usecase{
		productRepo: productRepo,
	}
}

func (r *usecase) InsertProduct(pctx context.Context) error {
	_ = r.productRepo.InsertProduct(pctx)
	return nil
}
