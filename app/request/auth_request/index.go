package auth_request

type Register_Request struct {
	Name     *string `form:"name" binding:"required"`
	Email    *string `form:"email" binding:"required"`
	Password *string `form:"password" binding:"required"`
	No_Telp  *string `form:"no_telp" binding:"required"`
	TokoName *string `form:"toko_name" binding:"required"`
}

type Login_Request struct {
	Email    *string `form:"email" binding:"required"`
	Password *string `form:"password" binding:"required"`
}
