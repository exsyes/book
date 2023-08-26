package web

type BorrowerCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
