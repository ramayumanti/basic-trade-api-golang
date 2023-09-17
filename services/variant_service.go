package services

import (
	"basictrade-api/models"
	"basictrade-api/repositories"
	"basictrade-api/requests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VariantService interface {
	CreateVariant(*gin.Context) (*models.Variant, error)
	GetAllVariant() (*[]models.Variant, error)
	GetVariantById(string) (*models.Variant, error)
	UpdateVariant(*gin.Context) (*models.Variant, error)
	DeleteVariant(string) (*models.Variant, error)
}

type variantServiceImpl struct {
	VariantRepository repositories.VariantRepository
	ProductService    ProductService
}

func NewVariantService(newVariantRepository repositories.VariantRepository, newProductService ProductService) VariantService {
	return &variantServiceImpl{
		VariantRepository: newVariantRepository,
		ProductService:    newProductService,
	}
}

func (service *variantServiceImpl) CreateVariant(ctx *gin.Context) (*models.Variant, error) {

	var variantReq requests.VariantRequest
	if err := ctx.ShouldBind(&variantReq); err != nil {
		return nil, err
	}

	product, err := service.ProductService.GetProductById(ctx.PostForm("product_id"))
	if err != nil {
		return nil, err
	}

	newUUID := uuid.New()

	newVariant := models.Variant{
		VariantName: variantReq.VariantName,
		Quantity:    variantReq.Quantity,
		UUID:        newUUID.String(),
		ProductID:   product.ID,
	}

	variant, err := service.VariantRepository.CreateVariant(&newVariant)
	if err != nil {
		return nil, err
	}

	return variant, err
}

func (service *variantServiceImpl) GetAllVariant() (*[]models.Variant, error) {
	variants, err := service.VariantRepository.GetAllVariant()

	if err != nil {
		return nil, err
	}

	return variants, err
}

func (service *variantServiceImpl) GetVariantById(id string) (*models.Variant, error) {
	variant, err := service.VariantRepository.GetVariantById(id)

	if err != nil {
		return nil, err
	}

	return variant, err
}

func (service *variantServiceImpl) UpdateVariant(ctx *gin.Context) (*models.Variant, error) {

	var newVariant requests.VariantUpdateRequest
	err := ctx.ShouldBind(&newVariant)
	if err != nil {
		return nil, err
	}

	existingVariant, err := service.GetVariantById(ctx.PostForm("variantUUID"))
	if err != nil {
		return nil, err
	}

	existingVariant.VariantName = newVariant.VariantName
	existingVariant.Quantity = newVariant.Quantity

	res, err := service.VariantRepository.UpdateVariant(existingVariant)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *variantServiceImpl) DeleteVariant(id string) (*models.Variant, error) {
	variant, err := service.VariantRepository.DeleteVariant(id)

	if err != nil {
		return nil, err
	}

	return variant, err
}
