package controllers

import (
	"basictrade-api/models"
	"basictrade-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetAllProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(newProductService services.ProductService) ProductController {
	return &productControllerImpl{
		ProductService: newProductService,
	}
}

func (controller *productControllerImpl) CreateProduct(ctx *gin.Context) {

	res, err := controller.ProductService.CreateProduct(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (controller *productControllerImpl) GetAllProduct(ctx *gin.Context) {

	var results = &[]models.Product{}

	results, err := controller.ProductService.GetAllProduct(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

func (controller *productControllerImpl) GetProductById(ctx *gin.Context) {

	productId := ctx.Param("productUUID")

	res, err := controller.ProductService.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (controller *productControllerImpl) UpdateProduct(ctx *gin.Context) {

	res, err := controller.ProductService.UpdateProduct(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (controller *productControllerImpl) DeleteProduct(ctx *gin.Context) {

	productId := ctx.Param("productUUID")

	res, err := controller.ProductService.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}
