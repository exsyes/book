package web

type BorrowedBookUpdateRequest struct {
	Id_borrowed   int `json:"id_Borrowed" validate:"required,number"`
	Id_book       int `json:"id_book" validate:"required,number"`
	Quantity_back int `json:"quantity_back" validate:"required,number"`
}
