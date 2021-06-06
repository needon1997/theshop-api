package forms

type CatBrandForm struct {
	CategoryID int32 `json:"category_id" binding:"required"`
	BrandID    int32 `json:"brand_id" binding:"required"`
}
