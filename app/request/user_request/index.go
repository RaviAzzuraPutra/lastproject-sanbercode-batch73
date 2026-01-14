package user_request

type User_Request struct {
	Name     *string `form:"name" binding:"required"`
	Email    *string `form:"email" binding:"required"`
	Password *string `form:"password" binding:"required"`
	No_Telp  *string `form:"no_telp" binding:"required"`
}
