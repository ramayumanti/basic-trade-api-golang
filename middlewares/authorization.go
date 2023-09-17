package middlewares

import (
	"basictrade-api/database"
	"basictrade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productUuid := ctx.Param("productUUID")

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminUuid := adminData["uuid"].(string)

		var adminRes = models.Admin{}
		res := db.Where(models.Admin{UUID: adminUuid}).First(&adminRes)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Admin data Not Found",
			})
			return
		}

		var productRes = models.Product{}
		res = db.Where(models.Product{UUID: productUuid}).Take(&productRes)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Product data Not Found",
			})
			return
		}

		if productRes.AdminID != adminRes.ID {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			panic("You are not allowed to access this data !!!!")
			// return
		}

		ctx.Next()
	}
}

func VariantCreateAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productUuid := ctx.PostForm("product_id")

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminUuid := adminData["uuid"].(string)

		var adminRes = models.Admin{}
		res := db.Where(models.Admin{UUID: adminUuid}).First(&adminRes)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Admin data Not Found",
			})
			return
		}

		var productRes = models.Product{}
		res = db.Where(models.Product{UUID: productUuid}).Take(&productRes)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Product data Not Found",
			})
			return
		}

		if productRes.AdminID != adminRes.ID {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to add variant to this product",
			})
			panic("You are not allowed to access this data !!!!")
			// return
		}

		ctx.Next()
	}
}

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		variantUuid := ctx.Param("variantUUID")
		var variantRes = models.Variant{}
		res := db.Where(models.Variant{UUID: variantUuid}).First(&variantRes)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Variant data Not Found",
			})
			return
		}

		var productRes = models.Product{}
		productId := variantRes.ProductID
		res = db.First(&productRes, productId)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Product data Not Found",
			})
			return
		}

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminUuid := adminData["uuid"].(string)

		var adminRes = models.Admin{}
		res = db.Where(models.Admin{UUID: adminUuid}).First(&adminRes)
		if res.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   res.Error,
				"message": "Admin data Not Found",
			})
			return
		}

		if productRes.AdminID != adminRes.ID {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			panic("You are not allowed to access this data !!!!")
			// return
		}

		ctx.Next()
	}
}
