package request

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required,min=2,max=255" json:"password"`
	Captcha  string `json:"captcha"`
}

type UpdateUserRequest struct {
	ID       uint   `validate:"required,min=2,max=100" json:"userID"`
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required,min=2,max=255" json:"password"`
	NFAID    *uint  `json:"nfaID"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Captcha  string `json:"captcha"`
}
