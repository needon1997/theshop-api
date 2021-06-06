package forms

type BannerForm struct {
	Image string `json:"image" binding:"required"`
	Url   string `json:"url" binding:"required,url"`
	Index int32  `json:"index" binding:"required,gte=1"`
}
