package repositories

import (
	"basictrade-api/database"
	"basictrade-api/models"
)

type ProductRepository interface {
	CreateProduct(*models.Product) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
	GetProductById(id string) (*models.Product, error)
	UpdateProduct(*models.Product) (*models.Product, error)
	DeleteProduct(string) (*models.Product, error)
}

type productRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}

func (repository *productRepositoryImpl) CreateProduct(ProductReq *models.Product) (*models.Product, error) {
	db := database.GetDB()
	err := db.Create(&ProductReq).Error
	if err != nil {
		return nil, err
	}

	return ProductReq, nil
}

func (repository *productRepositoryImpl) GetAllProduct() (*[]models.Product, error) {
	db := database.GetDB()
	var results = []models.Product{}

	res := db.Preload("Variants").Find(&results)
	if res.Error != nil {
		return nil, res.Error
	}

	return &results, nil
}

func (repository *productRepositoryImpl) GetProductById(id string) (*models.Product, error) {
	db := database.GetDB()
	var results = models.Product{}

	res := db.Where(models.Product{UUID: id}).Preload("Variants").First(&results)
	if res.Error != nil {
		return nil, res.Error
	}

	return &results, nil
}

func (repository *productRepositoryImpl) UpdateProduct(ProductReq *models.Product) (*models.Product, error) {
	db := database.GetDB()

	existingProduct := models.Product{}
	res := db.Preload("Variants").First(&existingProduct, ProductReq.ID)
	if res.Error != nil {
		return nil, res.Error
	}

	tx := db.Begin()

	existingProduct.Name = ProductReq.Name

	for i, newVariant := range ProductReq.Variants {
		if i < len(existingProduct.Variants) {
			existingProduct.Variants[i].VariantName = newVariant.VariantName
			existingProduct.Variants[i].Quantity = newVariant.Quantity
			if err := tx.Save(&existingProduct.Variants[i]).Error; err != nil {
				tx.Rollback()
				return nil, res.Error
			}
		} else {
			existingProduct.Variants = append(existingProduct.Variants, newVariant)
			if err := tx.Create(&existingProduct.Variants[i]).Error; err != nil {
				tx.Rollback()
				return nil, res.Error
			}
		}
	}

	if err := tx.Save(&existingProduct).Error; err != nil {
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()

	return &existingProduct, nil
}

func (repository *productRepositoryImpl) DeleteProduct(id string) (*models.Product, error) {
	db := database.GetDB()
	var product models.Product

	res := db.Where(models.Product{UUID: id}).Preload("Variants").First(&product)
	if res.Error != nil {
		return nil, res.Error
	}

	tx := db.Begin()

	if err := tx.Delete(&product.Variants).Error; err != nil {
		tx.Rollback()
		return nil, res.Error
	}

	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()

	return nil, nil
}
