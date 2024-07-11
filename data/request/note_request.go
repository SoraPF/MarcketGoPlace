package request

type CreateNoteRequest struct {
	Content string `validate:"required,min=2,max=100" json="content"`
}

type UpdateNoteRequest struct {
	Id      int    `validate:"required"`
	Content string `validate:"required,min=2,max=200" json="content"`
}
