package web

type BorrowerUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"Name" validate:"required"`
}
