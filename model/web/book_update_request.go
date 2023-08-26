package web

type BookUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Title    string `validate:"required,min=1,max=200" json:"title"`
	Quantity int    `validate:"required,min=1,max=200" json:"quantity"`
}
