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

func (u *productUsecase) DeleteProductAdmin(pctx context.Context, productId uint, userId uint) error {
	err := u.productRepo.DeleteProductAdmin(pctx, productId, userId)
	if err != nil {
		return err
	}
	return nil
}
