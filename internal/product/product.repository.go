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
		DeleteProduct(pctx context.Context, productId uint, userId uint) error
		UpdateProductAdmin(pctx context.Context, product *entity.Product, productId uint, userId uint) error
		DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error
		IsAdmin(ctx context.Context, db *gorm.DB, userId uint) (bool, error)
		IsOwner(ctx context.Context, db *gorm.DB, productId uint, userId uint) (bool, error)
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

func (r *productRepository) IsOwner(ctx context.Context, db *gorm.DB, productId uint, userId uint) (bool, error) {
	var product struct {
		UserID uint
	}
	if err := db.WithContext(ctx).Table("products").Select("user_id").Where("id = ?", productId).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("product not found")
		}
		return false, err
	}

	return product.UserID == userId, nil
}

// TODO: Find better place for this function since it's not related to product
func (r *productRepository) IsAdmin(ctx context.Context, db *gorm.DB, userId uint) (bool, error) {
	var user struct {
		Role string
	}
	if err := db.WithContext(ctx).Table("users").Select("role").Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("user not found")
		}
		return false, err
	}

	return user.Role == "admin", nil
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
	isOwner, err := r.IsOwner(pctx, r.Db, productId, userId)
	if err != nil {
		return err
	}

	if !isOwner {
		return errors.New("unauthorized: user does not have permission to update this product")
	}

	if err := r.productTable(pctx).Where("id = ? AND user_id = ?", productId, userId).Updates(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(pctx context.Context, productId uint, userId uint) error {
	isOwner, err := r.IsOwner(pctx, r.Db, productId, userId)
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
	isAdmin, err := r.IsAdmin(pctx, r.Db, userId)
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
	isAdmin, err := r.IsAdmin(pctx, r.Db, userId)
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
