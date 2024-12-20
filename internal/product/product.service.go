package product

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/pkg/jwt"
	"Pelter_backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type (
	productService struct {
		productUsecase ProductUsecase
	}
	ProductService interface {
		InsertProduct(ctx *fiber.Ctx) error
		GetProduct(ctx *fiber.Ctx) error
		GetProductIn(ctx *fiber.Ctx) error
		GetProductByID(ctx *fiber.Ctx) error
		UpdateProduct(ctx *fiber.Ctx) error
		DeleteProduct(ctx *fiber.Ctx) error
		UpdateProductAdmin(ctx *fiber.Ctx) error
		DeleteProductAdmin(ctx *fiber.Ctx) error
		UpdateProductVerificationStatus(ctx *fiber.Ctx) error
		GetProductByBuyerId(ctx *fiber.Ctx) error
		GetProductByUserId(ctx *fiber.Ctx) error
		UpdateProductIsSold(ctx *fiber.Ctx) error
	}
)

func NewProductService(productUsecase ProductUsecase) ProductService {
	return &productService{
		productUsecase: productUsecase,
	}
}

func (s *productService) InsertProduct(ctx *fiber.Ctx) error {
	var req dto.ProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: "Cannot get UserID from access_token context: " + err.Error(),
		})
	}

	product := entity.Product{
		UserID:         userId,
		Name:           req.Name,
		IsSold:         req.IsSold,
		Category:       req.Category,
		Subcategory:    req.Subcategory,
		Description:    req.Description,
		IsVerified:     req.IsVerified,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		VaccineBookURL: req.VaccineBookURL,
	}

	productRes, err := s.productUsecase.InsertProduct(ctx.UserContext(), &product)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Failed to insert product",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.HttpResponse{
		Result:  productRes,
		Success: true,
	})
}

func (s *productService) GetProduct(ctx *fiber.Ctx) error {
	products, err := s.productUsecase.GetProduct(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Failed to get product",
			Success: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  products,
		Success: true,
	})
}

func (s *productService) GetProductByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	product, err := s.productUsecase.GetProductByID(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  product,
		Success: true,
	})
}

func (s *productService) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	var req dto.ProductUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	updateProduct := entity.Product{
		UserID:         userId,
		Name:           req.Name,
		IsSold:         req.IsSold,
		Category:       req.Category,
		Subcategory:    req.Subcategory,
		Description:    req.Description,
		IsVerified:     req.IsVerified,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		VaccineBookURL: req.VaccineBookURL,
	}

	err = s.productUsecase.UpdateProduct(ctx.UserContext(), &updateProduct, id, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	// Get updated product and response
	product, err := s.productUsecase.GetProductByID(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  product,
		Success: true,
	})
}

func (s *productService) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	err = s.productUsecase.DeleteProduct(ctx.UserContext(), id, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  "deleted product",
		Success: true,
	})
}

func (s *productService) UpdateProductAdmin(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	var req dto.ProductUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	updateProduct := entity.Product{
		Name:           req.Name,
		IsSold:         req.IsSold,
		Category:       req.Category,
		Subcategory:    req.Subcategory,
		Description:    req.Description,
		IsVerified:     req.IsVerified,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		VaccineBookURL: req.VaccineBookURL,
	}

	err = s.productUsecase.UpdateProductAdmin(ctx.UserContext(), &updateProduct, id, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	// Get updated product and response
	product, err := s.productUsecase.GetProductByID(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  product,
		Success: true,
	})
}

func (s *productService) UpdateProductVerificationStatus(ctx *fiber.Ctx) error {
	productId, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	// userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
	// 		Error:   "Cannot get UserID from access_token context: " + err.Error(),
	// 		Success: false,
	// 	})
	// }

	var req dto.ProductUpdateVerificationStatus
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	err = s.productUsecase.UpdateVerificationStatus(ctx.UserContext(), productId, req.IsVerified)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	product, err := s.productUsecase.GetProductByID(ctx.UserContext(), productId)
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  product,
		Success: true,
	})
}

func (s *productService) DeleteProductAdmin(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	err = s.productUsecase.DeleteProductAdmin(ctx.UserContext(), id, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  "deleted product",
		Success: true,
	})
}

func (s *productService) GetProductByBuyerId(ctx *fiber.Ctx) error {

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	transactions, err := s.productUsecase.GetProductByBuyerId(ctx.UserContext(), userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  transactions,
		Success: true,
	})

}
func (s *productService) GetProductByUserId(ctx *fiber.Ctx) error {

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	products, err := s.productUsecase.GetProductByUserId(ctx.UserContext(), userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  products,
		Success: true,
	})

}

func (s *productService) UpdateProductIsSold(ctx *fiber.Ctx) error {

	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}

	productId, err := utils.ParseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	var req dto.UpdateProductIsSold
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	products, err := s.productUsecase.UpdateProductIsSold(ctx.UserContext(), productId, userId, req.IsSold)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  products,
		Success: true,
	})

}

func (s *productService) GetProductIn(ctx *fiber.Ctx) error {

	var req dto.QueryProductsIn
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	products, err := s.productUsecase.GetProductIn(ctx.UserContext(), req.ProductsId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  products,
		Success: true,
	})

}
