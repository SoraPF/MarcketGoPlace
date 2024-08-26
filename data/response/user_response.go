package response

type UserResponse struct {
	ID       uint `json:"id"`
	Name     string
	Email    string
	BirthDay string
	Address  string
	Phone    string
	NFAID    *uint `json:"nfaID"`
}

type NFAResponse struct {
	ID     int    `json:"NFAId"`
	QRcode string `json:"QRcode"`
}
