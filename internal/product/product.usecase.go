package product

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/entity"
	"context"
)

type (
	productUsecase struct {
		productRepo ProductRepository
	}
	ProductUsecase interface {
		InsertProduct(pctx context.Context, product *entity.Product) (dto.ProductResponse, error)
		GetProduct(pctx context.Context) ([]dto.ProductResponse, error)
		GetProductByID(pctx context.Context, productId uint) (dto.ProductResponse, error)
		UpdateProduct(pctx context.Context, product *entity.Product, productId uint, userId uint) error
		DeleteProduct(pctx context.Context, productId uint, userId uint) error
		UpdateProductAdmin(pctx context.Context, product *entity.Product, productId uint, userId uint) error
		DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error
		UpdateVerificationStatus(pctx context.Context, productId uint, isVerify bool) error
		GetProductByBuyerId(pctx context.Context, userId uint) ([]*entity.Transaction, error)
		GetProductByUserId(pctx context.Context, userId uint) ([]*dto.ProductWithUserResponse, error)
		UpdateProductIsSold(pctx context.Context, productId, userId uint, isSold bool) (dto.ProductResponse, error)
		GetProductIn(pctx context.Context, productIds []uint) ([]dto.ProductResponse, error)
	}
)

func NewProductUsecase(productRepo ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) InsertProduct(pctx context.Context, product *entity.Product) (dto.ProductResponse, error) {
	id, err := u.productRepo.InsertProduct(pctx, product)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	product.ID = id
	productResponse := product.ConvertToProductResponse()

	return productResponse, nil
}

func (u *productUsecase) GetProduct(pctx context.Context) ([]dto.ProductResponse, error) {
	products, err := u.productRepo.GetProduct(pctx)
	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, product.ConvertToProductResponse())
	}
	return productResponses, nil
}

func (u *productUsecase) GetProductIn(pctx context.Context, productIds []uint) ([]dto.ProductResponse, error) {
	products, err := u.productRepo.GetProductIn(pctx, productIds)
	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, product.ConvertToProductResponse())
	}
	return productResponses, nil
}

func (u *productUsecase) GetProductByID(pctx context.Context, productId uint) (dto.ProductResponse, error) {
	product, err := u.productRepo.GetProductByID(pctx, productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return product.ConvertToProductResponse(), nil
}

func (u *productUsecase) UpdateProduct(pctx context.Context, product *entity.Product, productId uint, userId uint) error {
	err := u.productRepo.UpdateProduct(pctx, product, productId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) DeleteProduct(pctx context.Context, productId uint, userId uint) error {
	err := u.productRepo.DeleteProduct(pctx, productId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) UpdateProductAdmin(pctx context.Context, product *entity.Product, productId uint, userId uint) error {
	err := u.productRepo.UpdateProductAdmin(pctx, product, productId, userId)
	if err != nil {
		return err
	}
	return nil
}
func (u *productUsecase) UpdateVerificationStatus(pctx context.Context, productId uint, isVerify bool) error {
	err := u.productRepo.UpdateVerificationStatus(pctx, productId, isVerify)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error {
	err := u.productRepo.DeleteProductAdmin(pctx, productId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) GetProductByBuyerId(pctx context.Context, userId uint) ([]*entity.Transaction, error) {
	products, err := u.productRepo.GetProductByBuyerId(pctx, userId)
	if err != nil {
		return []*entity.Transaction{}, err
	}

	// var productResponses []dto.ProductResponse
	// for _, product := range products {
	// 	productResponses = append(productResponses, product.ConvertToProductResponse())
	// }
	return products, nil

}
func (u *productUsecase) GetProductByUserId(pctx context.Context, userId uint) ([]*dto.ProductWithUserResponse, error) {
	products, err := u.productRepo.GetProductByUserId(pctx, userId)
	if err != nil {
		return []*dto.ProductWithUserResponse{}, err
	}

	var productResponses []*dto.ProductWithUserResponse
	for _, product := range products {
		productResponses = append(productResponses, &dto.ProductWithUserResponse{
			ID:          product.ID,
			UserID:      product.UserID,
			Name:        product.Name,
			IsSold:      product.IsSold,
			IsVerified:  product.IsVerified,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Description: product.Description,
			Category:    product.Category,
			Subcategory: product.Subcategory,
		})
	}
	return productResponses, nil
}

func (u *productUsecase) UpdateProductIsSold(pctx context.Context, productId, userId uint, isSold bool) (dto.ProductResponse, error) {
	err := u.productRepo.UpdateProductIsSoldById(pctx, productId, userId, isSold)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	product, err := u.productRepo.GetProductByID(pctx, userId)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return product.ConvertToProductResponse(), nil
}
