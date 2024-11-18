package product

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/utils"
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
		DeleteProduct(pctx context.Context, productId uint, userId uint) error
		UpdateProductAdmin(pctx context.Context, product *entity.Product, productId uint, userId uint) error
		DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error
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
	isOwner, err := utils.IsOwner(pctx, r.Db, userId, productId)
	if err != nil {
		return err
	}

	if !isOwner {
		fmt.Println("isOwner:", isOwner)
		return errors.New("unauthorized: user does not have permission to update this product")
	}

	if err := r.productTable(pctx).Where("id = ? AND user_id = ?", productId, userId).Updates(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(pctx context.Context, productId uint, userId uint) error {
	isOwner, err := utils.IsOwner(pctx, r.Db, productId, userId)
	if err != nil {
		return err
	}

	if !isOwner {
		return errors.New("unauthorized: user does not have permission to update this product")
	}

	if err := r.productTable(pctx).Where("id = ? AND user_id = ?", productId, userId).Delete(&entity.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateProductAdmin(pctx context.Context, product *entity.Product, productId uint, userId uint) error {
	isAdmin, err := utils.IsAdmin(pctx, r.Db, userId)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("unauthorized: user does not have permission to update this product")
	}

	if err := r.productTable(pctx).Where("id = ?", productId).Updates(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error {
	isAdmin, err := utils.IsAdmin(pctx, r.Db, userId)
	if err != nil {
		return err
	}

	if !isAdmin {
		return errors.New("unauthorized: user does not have permission to update this product")
	}

	if err := r.productTable(pctx).Where("id = ?", productId).Delete(&entity.Product{}).Error; err != nil {
		return err
	}

	return nil
}
