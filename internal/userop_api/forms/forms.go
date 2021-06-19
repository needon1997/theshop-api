package forms

type AddressForm struct {
	Province     string `form:"province" json:"province" binding:"required"`
	City         string `form:"city" json:"city" binding:"required"`
	Address      string `form:"address" json:"address" binding:"required"`
	SignerName   string `form:"signer_name" json:"signer_name" binding:"required"`
	SignerMobile string `form:"signer_mobile" json:"signer_mobile" binding:"required"`
}
type MessageForm struct {
	MessageType string `form:"type" json:"type" binding:"required,oneof='note', 'complaint', 'inquiry', 'customer service', 'quote'"`
	Subject     string `form:"subject" json:"subject" binding:"required"`
	Message     string `form:"message" json:"message" binding:"required"`
	File        string `form:"file" json:"file" binding:"required"`
}
type UserFavForm struct {
	GoodsId int32 `form:"goods" json:"goods" binding:"required"`
}
