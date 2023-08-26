package web

type BorrowedBookCreateRequest struct {
	Id_borrower       int `json:"id_Borrower" validate:"required,number"`
	Id_book           int `json:"id_book" validate:"required,number"`
	Quantity_borrowed int `json:"quantity_borrowed" validate:"required,number"`
}
