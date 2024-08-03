package request

type CreateObjRequest struct {
	IdVendeur  int      `validate:"required" json:"id_vendeur"`
	Title      string   `validate:"required,min=2,max=100" json:"title"`
	Price      int      `validate:"required" json:"price"`
	Desc       string   `validate:"required,min=2,max=255" json:"desc"`
	StatusID   int      `validate:"required" json:"status_id"`
	CategoryID int      `validate:"required" json:"category_id"`
	Tags       []int    `validate:"required" json:"tags"`
	Img        []string `json:"images"`
}

type UpdateObjRequest struct {
	ID         int    `validate:"required" json:"objId"`
	IdVendeur  int    `json:"id_vendeur"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Desc       string `json:"desc"`
	StatusID   int    `json:"status_id"`
	CategoryID int    `json:"category_id"`
	Tags       []int  `json:"tags"`
}
