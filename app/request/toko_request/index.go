package toko_request

type Toko_Request struct {
	Name    *string `form:"name" binding:"required"`
	Address *string `form:"address" binding:"required"`
}
