package request

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required,min=2,max=255" json:"password"`
	Captcha  string `json:"captcha"`
}

type UpdateUserRequest struct {
	ID       int    `validate:"min=2,max=100" json:"userID"`
	Username string `validate:"min=2,max=100" json:"username"`
	Email    string `json:"email"`
	Password string `validate:"min=2,max=255" json:"password"`
	BirthDay string `validate:"min=5,max=30" json:"BirthDay"`
	Phone    string `validate:"min=10,max=15" json:"Phone"`
	Address  string `validate:"min=5,max=100" json:"Address"`
	NFAID    *uint  `json:"nfaID"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Captcha  string `json:"captcha"`
}
