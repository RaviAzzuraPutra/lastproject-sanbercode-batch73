package trx_request

type Trx_Log_Request struct {
	Qty  *int    `form:"qty" binding:"required,gt=0"`
	Type *string `form:"type" binding:"required,oneof=in out"`
}
