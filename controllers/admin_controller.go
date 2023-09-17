package controllers

import (
	"basictrade-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController interface {
	RegisterAdmin(ctx *gin.Context)
	LoginAdmin(ctx *gin.Context)
}

type adminControllerImpl struct {
	AdminService services.AdminService
}

func NewAdminController(newAdminService services.AdminService) AdminController {
	return &adminControllerImpl{
		AdminService: newAdminService,
	}
}

func (controller *adminControllerImpl) LoginAdmin(ctx *gin.Context) {

	token, err := controller.AdminService.LoginAdmin(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})
}

func (controller *adminControllerImpl) RegisterAdmin(ctx *gin.Context) {

	res, err := controller.AdminService.RegisterAdmin(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}
