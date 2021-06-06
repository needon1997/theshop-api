package forms

type BrandForm struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
	Logo string `json:"logo" binding:"required,url,max=256"`
}
