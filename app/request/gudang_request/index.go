package gudang_request

type Gudang_Request struct {
	Name    *string `form:"name" binding:"required"`
	Address *string `form:"address" binding:"required"`
}
