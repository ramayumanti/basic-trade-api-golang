package repositories

import (
	"basictrade-api/database"
	"basictrade-api/models"
)

type VariantRepository interface {
	CreateVariant(*models.Variant) (*models.Variant, error)
	GetAllVariant(offset int, limit int, search string) (*[]models.Variant, error)
	GetVariantById(string) (*models.Variant, error)
	UpdateVariant(*models.Variant) (*models.Variant, error)
	DeleteVariant(string) (*models.Variant, error)
}

type variantRepositoryImpl struct{}

func NewVariantRepository() VariantRepository {
	return &variantRepositoryImpl{}
}

func (repository *variantRepositoryImpl) CreateVariant(VariantReq *models.Variant) (*models.Variant, error) {
	db := database.GetDB()

	err := db.Create(&VariantReq).Error

	if err != nil {
		return nil, err
	}

	return VariantReq, nil
}

func (repository *variantRepositoryImpl) GetAllVariant(offset int, limit int, search string) (*[]models.Variant, error) {
	db := database.GetDB()
	var results = []models.Variant{}

	if search != "" {
		res := db.Offset(offset).Limit(limit).Where("variant_name LIKE ? ", "%"+search+"%").Find(&results)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		res := db.Offset(offset).Limit(limit).Find(&results)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	return &results, nil
}

func (repository *variantRepositoryImpl) GetVariantById(id string) (*models.Variant, error) {
	db := database.GetDB()
	var results = models.Variant{}

	res := db.Where(models.Variant{UUID: id}).Take(&results)
	if res.Error != nil {
		return nil, res.Error
	}

	return &results, nil
}

func (repository *variantRepositoryImpl) UpdateVariant(VariantReq *models.Variant) (*models.Variant, error) {
	db := database.GetDB()

	existingVariant := models.Variant{}
	res := db.First(&existingVariant, VariantReq.ID)
	if res.Error != nil {
		return nil, res.Error
	}

	tx := db.Begin()

	existingVariant.VariantName = VariantReq.VariantName
	existingVariant.Quantity = VariantReq.Quantity

	if err := tx.Save(&existingVariant).Error; err != nil {
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()

	return &existingVariant, nil
}

func (repository *variantRepositoryImpl) DeleteVariant(id string) (*models.Variant, error) {
	db := database.GetDB()
	var variant models.Variant

	res := db.Where(models.Variant{UUID: id}).Take(&variant)
	if res.Error != nil {
		return nil, res.Error
	}

	tx := db.Begin()

	if err := tx.Delete(&variant).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return nil, nil
}
