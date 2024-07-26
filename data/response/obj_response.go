package response

type ObjResponse struct {
	IdVendeur int              `json:"id_vendeur"`
	Title     string           `json:"title"`
	Price     int              `json:"price"`
	Desc      string           `json:"desc"`
	Status    StatusResponse   `json:"status"`
	Category  CategoryResponse `json:"category"`
	Tags      []TagResponse    `json:"tags"`
}

type StatusResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type CategoryResponse struct {
	ID    uint
	Title string
	Img   string
}

type TagResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
