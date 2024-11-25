package product

import (
	"context"
	"errors"
	"fmt"

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
		UpdateVerificationStatus(pctx context.Context, productId uint, isVerified bool) error
		DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error
		IsAdmin(ctx context.Context, db *gorm.DB, userId uint) (bool, error)
		IsOwner(ctx context.Context, db *gorm.DB, productId uint, userId uint) (bool, error)
		GetProductByBuyerId(pctx context.Context, userId uint) ([]*entity.Transaction, error)
		GetProductByUserId(pctx context.Context, userId uint) ([]entity.Product, error)
		UpdateProductIsSoldById(pctx context.Context, productId, userId uint, is_sold bool) error
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

	result := r.productTable(pctx).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, role")
		}).
		Find(&products)

	if result.Error != nil {
		return []entity.Product{}, fmt.Errorf("failed to get products: %w", result.Error)
	}

	return products, nil
}

func (r *productRepository) GetProductByID(pctx context.Context, productId uint) (entity.Product, error) {
	var product entity.Product
	if err := r.productTable(pctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, profile_url, phone_number")
	}).Where("id = ?", productId).First(&product).Error; err != nil {
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

func (r *productRepository) UpdateVerificationStatus(pctx context.Context, productId uint, isVerified bool) error {
	// TODO: if have time implement isAdmin
	result := r.productTable(pctx).
		Model(&entity.Product{}).
		Where("id = ?", productId).
		Update("is_verified", isVerified)

	if result.Error != nil {
		return fmt.Errorf("failed to update verification status: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product with id %d not found", productId)
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

func (r *productRepository) GetProductByBuyerId(pctx context.Context, userId uint) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction

	result := r.Db.WithContext(pctx).
		Table("transactions").
		Select("transactions.*, users.name as buyer_name, users.profile_url, users.phone_number, products.*").
		Joins("LEFT JOIN users ON users.id = transactions.buyer_id").
		Joins("LEFT JOIN products ON products.id = transactions.product_id").
		Where("transactions.buyer_id = ?", userId).
		Preload("Buyer").
		Preload("Product").
		Find(&transactions)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to find transactions: %w", result.Error)
	}

	if len(transactions) == 0 {
		return nil, nil
	}

	return transactions, nil
}

func (r *productRepository) GetProductByUserId(pctx context.Context, userId uint) ([]entity.Product, error) {
	var products []entity.Product

	result := r.productTable(pctx).
		Where("user_id = ?", userId).
		Find(&products)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to find products: %w", result.Error)
	}

	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}

func (r *productRepository) UpdateProductIsSoldById(pctx context.Context, productId, userId uint, is_sold bool) error {
	query := r.productTable(pctx).
		Where("id = ? AND user_id = ?", productId, userId).
		Update("is_sold", is_sold)

	if query.Error != nil {
		return fmt.Errorf("failed to update product: %w", query.Error)
	}

	if query.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
