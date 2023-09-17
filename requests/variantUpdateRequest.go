package requests

type VariantUpdateRequest struct {
	VariantName string `form:"variant_name" binding:"required"`
	Quantity    int    `form:"quantity" binding:"required"`
}
