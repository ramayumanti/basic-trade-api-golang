package repositories

import (
	"basictrade-api/database"
	"basictrade-api/models"
)

type AdminRepository interface {
	RegisterAdmin(*models.Admin) (*models.Admin, error)
	SearchAdminByEmail(email string) (*models.Admin, error)
}

type adminRepositoryImpl struct{}

func NewAdminRepository() AdminRepository {
	return &adminRepositoryImpl{}
}

func (repository *adminRepositoryImpl) RegisterAdmin(AdminReq *models.Admin) (*models.Admin, error) {

	db := database.GetDB()

	err := db.Create(&AdminReq).Error

	if err != nil {
		return nil, err
	}

	return AdminReq, nil
}

func (*adminRepositoryImpl) SearchAdminByEmail(email string) (*models.Admin, error) {

	db := database.GetDB()

	AdminRes := models.Admin{}
	err := db.Debug().Where("email = ?", email).Take(&AdminRes).Error
	if err != nil {
		return nil, err
	}

	return &AdminRes, nil
}
