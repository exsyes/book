package web

type BorrowedBookViewResponse struct {
	Id_borrowed       int    `json:"id_Borrowed"`
	Borrower_name     string `json:"borrower_name"`
	Book_title        string `json:"book_title"`
	Quantity_borrowed int    `json:"quantity_borrowed"`
	Quantity_back     int    `json:"quantity_back"`
	Date_borrowed     string `json:"date"`
}
