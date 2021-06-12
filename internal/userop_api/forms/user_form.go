package forms

type EmailPasswordLoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type EmailForm struct {
	Email string `json:"email" binding:"required,email"`
}
type RegisterForm struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}
