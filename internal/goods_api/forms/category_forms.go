package forms

type CategoryForm struct {
	Name   string `json:"name" binding:"required,min=3,max=20"`
	Parent int32  `json:"parent" binding:"required,oneof= 0 1 2"`
	Level  int32  `json:"level" binding:"required,oneof=1 2 3"`
	IsTab  *bool  `json:"is_tab" binding:"required"`
}

type CategoryUpdateForm struct {
	Name  string `json:"name" binding:"required,min=3,max=20"`
	IsTab *bool  `json:"is_tab" binding:"required"`
}
