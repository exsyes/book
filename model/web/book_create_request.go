package web

type BookCreateRequest struct {
	Title    string `validate:"required,min=1,max=200" json:"title"`
	Quantity int    `validate:"required" json:"quantity"`
}
