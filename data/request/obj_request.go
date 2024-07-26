package request

type CreateObjRequest struct {
	IdVendeur  int      `validate:"required" json:"id_vendeur"`
	Title      string   `validate:"required,min=2,max=100" json:"title"`
	Price      int      `validate:"required" json:"price"`
	Desc       string   `validate:"required,min=2,max=255" json:"desc"`
	StatusID   uint     `validate:"required" json:"status_id"`
	CategoryID uint     `validate:"required" json:"category_id"`
	Tags       []uint   `validate:"required" json:"tags"`
	Img        []string `validate:"required" json:"images"`
}

type UpdateObjRequest struct {
	ID         uint   `validate:"required" json:"objId"`
	IdVendeur  int    `validate:"required" json:"id_vendeur"`
	Title      string `validate:"required,min=2,max=100" json:"title"`
	Price      int    `validate:"required" json:"price"`
	Desc       string `validate:"required,min=2,max=255" json:"desc"`
	StatusID   uint   `validate:"required" json:"status_id"`
	CategoryID uint   `validate:"required" json:"category_id"`
	Tags       []uint `validate:"required" json:"tags"`
}
