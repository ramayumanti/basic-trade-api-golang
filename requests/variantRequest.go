package requests

type VariantRequest struct {
	VariantName string `form:"variant_name" binding:"required"`
	Quantity    int    `form:"quantity" binding:"required"`
	ProductID   string `form:"product_id" binding:"required"`
}
