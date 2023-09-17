package routers

import (
	"basictrade-api/controllers"
	"basictrade-api/database"
	"basictrade-api/middlewares"
	"basictrade-api/repositories"
	"basictrade-api/services"

	"github.com/gin-gonic/gin"
)

func init() {
	database.StartDB()
}

func StartServer() *gin.Engine {
	router := gin.Default()

	adminRepository := repositories.NewAdminRepository()
	adminService := services.NewAdminService(adminRepository)
	adminController := controllers.NewAdminController(adminService)

	authRoute := router.Group("/auth")
	{
		authRoute.POST("/register", adminController.RegisterAdmin)
		authRoute.POST("/login", adminController.LoginAdmin)
	}

	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository, adminRepository)
	productController := controllers.NewProductController(productService)

	variantRepository := repositories.NewVariantRepository()
	variantService := services.NewVariantService(variantRepository, productService)
	variantController := controllers.NewVariantController(variantService)

	productRoute := router.Group("/products")
	{
		productRoute.GET("/:productUUID", productController.GetProductById)
		productRoute.GET("/", productController.GetAllProduct)
		productRoute.GET("variants/:variantUUID", variantController.GetVariantById)
		productRoute.GET("variants/", variantController.GetAllVariant)

		productRoute.Use(middlewares.Authentication())
		productRoute.POST("/", productController.CreateProduct)
		productRoute.PUT("/:productUUID", middlewares.ProductAuthorization(), productController.UpdateProduct)
		productRoute.DELETE("/:productUUID", middlewares.ProductAuthorization(), productController.DeleteProduct)
		productRoute.POST("variants/", middlewares.VariantCreateAuthorization(), variantController.CreateVariant)
		productRoute.PUT("variants/:variantUUID", middlewares.VariantAuthorization(), variantController.UpdateVariant)
		productRoute.DELETE("variants/:variantUUID", middlewares.VariantAuthorization(), variantController.DeleteVariant)
	}

	return router
}
