package response

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type NFAResponse struct {
	ID     int    `json:"NFAId"`
	QRcode string `json:"QRcode"`
}
