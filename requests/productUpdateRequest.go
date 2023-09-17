package requests

type ProductUpdateRequest struct {
	Name string `form:"name" binding:"required"`
}
