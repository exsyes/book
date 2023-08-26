package domain

import "time"

type BorrowedBooks struct {
	Id_borrowed       int
	Id_book           int
	Id_borrower       int
	Quantity_borrowed int
	Quantity_back     int
	Date_borrowed     time.Time
}

type BorrowedBooksView struct {
	Id_borrowed       int
	Borrower_name     string
	Book_title        string
	Quantity_borrowed int
	Quantity_back     int
	Date_borrowed     string
}
