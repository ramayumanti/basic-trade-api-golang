package services

import (
	"basictrade-api/helpers"
	"basictrade-api/models"
	"basictrade-api/repositories"
	"basictrade-api/requests"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(*gin.Context) (*models.Product, error)
	GetAllProduct(*gin.Context) (*[]models.Product, error)
	GetProductById(id string) (*models.Product, error)
	UpdateProduct(*gin.Context) (*models.Product, error)
	DeleteProduct(string) (*models.Product, error)
}

type productServiceImpl struct {
	ProductRepository repositories.ProductRepository
	AdminRepository   repositories.AdminRepository
}

func NewProductService(newProductRepository repositories.ProductRepository, newAdminRepository repositories.AdminRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: newProductRepository,
		AdminRepository:   newAdminRepository,
	}
}

func (service *productServiceImpl) CreateProduct(ctx *gin.Context) (*models.Product, error) {

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	existingAdmin, err := service.AdminRepository.SearchAdminByEmail(adminData["email"].(string))
	if err != nil {
		return nil, err
	}

	var productReq requests.ProductRequest
	if err := ctx.ShouldBind(&productReq); err != nil {
		return nil, err
	}

	fileName := helpers.RemoveExtension(productReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		return nil, err
	}

	newUUID := uuid.New()

	newProduct := models.Product{
		Name:     productReq.Name,
		ImageURL: uploadResult,
		UUID:     newUUID.String(),
		AdminID:  existingAdmin.ID,
	}

	productRes, err := service.ProductRepository.CreateProduct(&newProduct)
	if err != nil {
		return nil, err
	}

	return productRes, nil
}

func (service *productServiceImpl) GetAllProduct(ctx *gin.Context) (*[]models.Product, error) {

	var offset, limit int
	var err error

	if ctx.Query("offset") != "" {
		offset, err = strconv.Atoi(ctx.Query("offset"))
		if err != nil {
			return nil, err
		}
	} else {
		offset = 0
	}

	if ctx.Query("limit") != "" {
		limit, err = strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			return nil, err
		}
	} else {
		limit = 50
	}

	search := ctx.Query("search")

	products, err := service.ProductRepository.GetAllProduct(offset, limit, search)
	if err != nil {
		return nil, err
	}

	return products, err
}

func (service *productServiceImpl) GetProductById(id string) (*models.Product, error) {
	product, err := service.ProductRepository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, err
}

func (service *productServiceImpl) UpdateProduct(ctx *gin.Context) (*models.Product, error) {

	var newProduct requests.ProductUpdateRequest
	err := ctx.ShouldBind(&newProduct)
	if err != nil {
		return nil, err
	}

	existingProduct, err := service.GetProductById(ctx.Param("productUUID"))
	if err != nil {
		return nil, err
	}

	existingProduct.Name = newProduct.Name

	res, err := service.ProductRepository.UpdateProduct(existingProduct)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *productServiceImpl) DeleteProduct(id string) (*models.Product, error) {
	product, err := service.ProductRepository.DeleteProduct(id)

	if err != nil {
		return nil, err
	}

	return product, err
}
