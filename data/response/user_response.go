package response

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	NFAID    *uint  `json:"nfaID"`
}

type NFAResponse struct {
	ID     int    `json:"NFAId"`
	QRcode string `json:"QRcode"`
}
