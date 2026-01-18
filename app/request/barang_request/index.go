package barang_request

type Barang_Request struct {
	Name           *string `form:"name" binding:"required"`
	Sku            *string `form:"sku" binding:"required"`
	Image_url      *string
	Stock          *int    `form:"stock" binding:"required"`
	Safety_stock   *int    `form:"safety_stock" binding:"required"`
	Lead_time_days *int    `form:"lead_time_days" binding:"required"`
	Category_id    *string `form:"category_id" binding:"required"`
}
