package helper

import (
	"buku/model/domain"
	"buku/model/web"
)

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		book.Id,
		book.Title,
		book.Quantity,
	}
}
func ToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}
	return bookResponses
}

func ToBorrowerResponse(borrower domain.Borrower) web.BorrowerResponse {
	return web.BorrowerResponse{
		Id:   borrower.Id,
		Name: borrower.Name,
	}
}

func ToBorrowerResponses(borrowers []domain.Borrower) []web.BorrowerResponse {
	var borrowerResponses []web.BorrowerResponse
	for _, borrower := range borrowers {
		borrowerResponses = append(borrowerResponses, ToBorrowerResponse(borrower))
	}
	return borrowerResponses
}

func ToBorrowedBooksResponse(borrowedBook domain.BorrowedBooks) web.BorrowedBookResponse {
	return web.BorrowedBookResponse{
		Id_book:           borrowedBook.Id_book,
		Id_borrower:       borrowedBook.Id_borrower,
		Id_borrowed:       borrowedBook.Id_borrowed,
		Quantity_borrowed: borrowedBook.Quantity_borrowed,
		Quantity_back:     borrowedBook.Quantity_back,
		Date_borrowed:     borrowedBook.Date_borrowed,
	}
}

func ToBorrowedBooksResponses(borrowedBooks []domain.BorrowedBooks) []web.BorrowedBookResponse {
	var borrowedBooksResponses []web.BorrowedBookResponse
	for _, borrowed := range borrowedBooks {
		borrowedBooksResponses = append(borrowedBooksResponses, ToBorrowedBooksResponse(borrowed))
	}
	return borrowedBooksResponses
}

func ToBorrowedBooksViewResponse(borrowedBook domain.BorrowedBooksView) web.BorrowedBookViewResponse {
	return web.BorrowedBookViewResponse{
		Id_borrowed:       borrowedBook.Id_borrowed,
		Borrower_name:     borrowedBook.Borrower_name,
		Book_title:        borrowedBook.Book_title,
		Quantity_borrowed: borrowedBook.Quantity_borrowed,
		Quantity_back:     borrowedBook.Quantity_back,
		Date_borrowed:     borrowedBook.Date_borrowed,
	}
}

func ToBorrowedBooksViewResponses(borrowedBooks []domain.BorrowedBooksView) []web.BorrowedBookViewResponse {
	var borrowedBooksViewResponses []web.BorrowedBookViewResponse
	for _, borrowedBook := range borrowedBooks {
		borrowedBooksViewResponses = append(borrowedBooksViewResponses, ToBorrowedBooksViewResponse(borrowedBook))
	}
	return borrowedBooksViewResponses
}
