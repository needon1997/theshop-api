package forms

type CartItemCreateForm struct {
	GoodsId int32 `json:"goods" binding:"required"`
	Nums    int32 `json:"nums" binding:"required,gte=1"`
}
type CartItemUpdateForm struct {
	GoodsId int32 `json:"goods" binding:"required"`
	Nums    int32 `json:"nums" binding:"required,gte=1"`
	Checked *bool `json:"checked" binding:"required"`
}
