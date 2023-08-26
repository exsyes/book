package web

type BookUpdateQtyRequest struct {
	Id       int `validate:"required"`
	Quantity int `validate:"required"json:"quantity"`
}
