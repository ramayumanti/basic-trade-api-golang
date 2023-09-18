package controllers

import (
	"basictrade-api/models"
	"basictrade-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VariantController interface {
	CreateVariant(ctx *gin.Context)
	GetAllVariant(ctx *gin.Context)
	GetVariantById(ctx *gin.Context)
	UpdateVariant(ctx *gin.Context)
	DeleteVariant(ctx *gin.Context)
}

type variantControllerImpl struct {
	VariantService services.VariantService
}

func NewVariantController(newVariantService services.VariantService) VariantController {
	return &variantControllerImpl{
		VariantService: newVariantService,
	}
}

func (controller *variantControllerImpl) CreateVariant(ctx *gin.Context) {

	res, err := controller.VariantService.CreateVariant(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (controller *variantControllerImpl) GetAllVariant(ctx *gin.Context) {

	var results = &[]models.Variant{}

	results, err := controller.VariantService.GetAllVariant(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Variant record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

func (controller *variantControllerImpl) GetVariantById(ctx *gin.Context) {

	variantId := ctx.Param("variantUUID")

	res, err := controller.VariantService.GetVariantById(variantId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Variant record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (controller *variantControllerImpl) UpdateVariant(ctx *gin.Context) {

	res, err := controller.VariantService.UpdateVariant(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (controller *variantControllerImpl) DeleteVariant(ctx *gin.Context) {

	variantId := ctx.Param("variantUUID")

	res, err := controller.VariantService.DeleteVariant(variantId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}
