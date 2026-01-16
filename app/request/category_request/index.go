package category_request

type Category_Request struct {
	Name        *string `form:"name" binding:"required"`
	Description *string `form:"description" binding:"required"`
}
