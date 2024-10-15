package product

import (
	"context"

	"gorm.io/gorm"
)

type (
	repository struct {
		Db *gorm.DB
	}
	ProductRepository interface {
		InsertProduct(pctx context.Context) error
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &repository{
		Db: db,
	}
}

func (r *repository) InsertProduct(pctx context.Context) error {
	return nil
}
