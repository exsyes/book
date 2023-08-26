package web

import "time"

type BorrowedBookResponse struct {
	Id_borrowed       int       `json:"id_Borrowed"`
	Id_borrower       int       `json:"id_Borrower"`
	Id_book           int       `json:"id_book"`
	Quantity_borrowed int       `json:"quantity_borrowed"`
	Quantity_back     int       `json:"quantity_back"`
	Date_borrowed     time.Time `json:"date"`
}
