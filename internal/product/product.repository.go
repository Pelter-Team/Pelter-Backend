package product

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"Pelter_backend/internal/entity"
)

type (
	productRepository struct {
		Db *gorm.DB
	}
	ProductRepository interface {
		GetProduct(pctx context.Context) ([]entity.Product, error)
		GetProductByID(pctx context.Context, productId uint) (entity.Product, error)
		InsertProduct(pctx context.Context, product *entity.Product) (uint, error)
		UpdateProduct(pctx context.Context, product *entity.Product, productId uint, userId uint) error
		DeleteProduct(pctx context.Context, userId uint, productId uint) error
	}
)

func (r *productRepository) productTable(pctx context.Context) *gorm.DB {
	return r.Db.Table("products").WithContext(pctx)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		Db: db,
	}
}

func (r *productRepository) InsertProduct(pctx context.Context, product *entity.Product) (uint, error) {
	if err := r.productTable(pctx).Create(product).Error; err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (r *productRepository) GetProduct(pctx context.Context) ([]entity.Product, error) {
	var products []entity.Product
	if err := r.productTable(pctx).Find(&products).Error; err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}

func (r *productRepository) GetProductByID(pctx context.Context, productId uint) (entity.Product, error) {
	var product entity.Product
	if err := r.productTable(pctx).Where("id = ?", productId).First(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (r *productRepository) UpdateProduct(pctx context.Context, product *entity.Product, productId uint, userId uint) error {
	query := r.productTable(pctx).Where("id = ? AND user_id = ?", productId, userId).Updates(product)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return errors.New("you don't have permission to edit this product")
	}

	return nil
}

func (r *productRepository) DeleteProduct(pctx context.Context, userId uint, productId uint) error {
	query := r.productTable(pctx).Where("id = ? AND user_id = ?", productId, userId).Delete(&entity.Product{})
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return errors.New("you don't have permission to delete this product")
	}

	return nil
}
