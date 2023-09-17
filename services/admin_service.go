package services

import (
	"basictrade-api/helpers"
	"basictrade-api/models"
	"basictrade-api/repositories"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminService interface {
	RegisterAdmin(ctx *gin.Context) (*models.Admin, error)
	LoginAdmin(ctx *gin.Context) (string, error)
}

type adminServiceImpl struct {
	AdminRepository repositories.AdminRepository
}

func NewAdminService(newAdminRepository repositories.AdminRepository) AdminService {
	return &adminServiceImpl{
		AdminRepository: newAdminRepository,
	}
}

func (service *adminServiceImpl) RegisterAdmin(ctx *gin.Context) (*models.Admin, error) {

	Admin := models.Admin{}
	if err := ctx.ShouldBind(&Admin); err != nil {
		return nil, err
	}

	newUUID := uuid.New()
	Admin.UUID = newUUID.String()

	adminRes, err := service.AdminRepository.RegisterAdmin(&Admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	return adminRes, nil
}

func (service *adminServiceImpl) LoginAdmin(ctx *gin.Context) (string, error) {

	contentType := helpers.GetContentType(ctx)
	loginData := models.Admin{}

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&loginData)
	} else {
		ctx.ShouldBind(&loginData)
	}

	adminRes, err := service.AdminRepository.SearchAdminByEmail(loginData.Email)
	if err != nil {
		err := errors.New("invalid email")
		return "", err
	}

	comparePass := helpers.ComparePass([]byte(adminRes.Password), []byte(loginData.Password))
	if !comparePass {
		err := errors.New("invalid password")
		return "", err
	}

	token := helpers.GenerateToken(adminRes.UUID, adminRes.Email)

	return token, nil
}
