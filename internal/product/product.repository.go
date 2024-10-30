package product

import (
	"context"

	"gorm.io/gorm"
)

type (
	productRepository struct {
		Db *gorm.DB
	}
	ProductRepository interface {
		InsertProduct(pctx context.Context) error
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		Db: db,
	}
}

func (r *productRepository) InsertProduct(pctx context.Context) error {
	return nil
}

func (r *productRepository) DeleteProduct(pctx context.Context) error {
	return nil
}
