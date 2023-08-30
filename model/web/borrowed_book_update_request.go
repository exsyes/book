package web

type BorrowedBookUpdateRequest struct {
	Id_borrowed   int `json:"id_Borrowed" validate:"required"`
	Id_book       int `json:"id_book" validate:"required"`
	Quantity_back int `json:"quantity_back" validate:"required"`
}
