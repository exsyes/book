package repository

import (
	"buku/helper"
	"buku/model/domain"
	"context"
	"database/sql"
)

type BorrowedBookRepositoryImpl struct {
}

func NewBorrowedBookRepository() BorrowedBookRepository {
	return &BorrowedBookRepositoryImpl{}
}

func (repository *BorrowedBookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, BorrowedBook domain.BorrowedBooks) domain.BorrowedBooks {
	SQL := `INSERT INTO borrowed_books(id_book, id_borrower, quantity_borrowed, quantity_back, DATE_BORROWED)
VALUE(?, ?, ?, 0, NOW())`

	result, err := tx.ExecContext(ctx, SQL, BorrowedBook.Id_book, BorrowedBook.Id_borrower, BorrowedBook.Quantity_borrowed)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	BorrowedBook.Id_borrowed = int(id)

	return BorrowedBook
}

func (repository *BorrowedBookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, BorrowedBook domain.BorrowedBooksView) domain.BorrowedBooksView {
	SQL := `UPDATE borrowed_books SET quantity_back = ?, quantity_borrowed = ? WHERE id_borrowed = ?`
	_, err := tx.ExecContext(ctx, SQL, BorrowedBook.Quantity_back, BorrowedBook.Quantity_borrowed, BorrowedBook.Id_borrowed)
	helper.PanicIfError(err)
	return BorrowedBook
}

func (repository *BorrowedBookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, BorrowedBook domain.BorrowedBooks) {
	SQL := `DELETE FROM borrowed_books WHERE id_borrowed = ?`
	_, err := tx.ExecContext(ctx, SQL, BorrowedBook.Id_borrowed)
	helper.PanicIfError(err)
}

func (repository *BorrowedBookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, BorrowedBookId int) (domain.BorrowedBooks, error) {
	SQL := `SELECT id_borrowed, id_book, id_borrower, quantity_borrowed, quantity_back, date_borrowed FROM borrowed_books WHERE id_borrowed = ?`
	rows, err := tx.QueryContext(ctx, SQL, BorrowedBookId)
	helper.PanicIfError(err)
	rows.Close()

	BorrowedBook := domain.BorrowedBooks{}
	if rows.Next() {
		err := rows.Scan(&BorrowedBook.Id_borrowed, &BorrowedBook.Id_borrower, &BorrowedBook.Id_book, &BorrowedBook.Quantity_borrowed, &BorrowedBook.Quantity_back, &BorrowedBook.Date_borrowed)
		helper.PanicIfError(err)
		return BorrowedBook, nil
	} else {
		return BorrowedBook, err
	}
}

func (repository *BorrowedBookRepositoryImpl) FindByIdView(ctx context.Context, tx *sql.Tx, BorrowedBookId int) (domain.BorrowedBooksView, error) {
	SQL := `SELECT id_borrowed, borrower.name as borrower_name, book.title as book_title, quantity_borrowed, quantity_back, date_borrowed FROM borrowed_books
		INNER JOIN borrower ON borrower.id = borrowed_books.id_borrower
		INNER JOIN book ON book.id = borrowed_books.id_book
		WHERE id_borrowed = ?;`

	rows, err := tx.QueryContext(ctx, SQL, BorrowedBookId)
	helper.PanicIfError(err)
	defer rows.Close()

	BorrowedBook := domain.BorrowedBooksView{}
	if rows.Next() {
		err := rows.Scan(&BorrowedBook.Id_borrowed,
			&BorrowedBook.Borrower_name,
			&BorrowedBook.Book_title,
			&BorrowedBook.Quantity_borrowed,
			&BorrowedBook.Quantity_back,
			&BorrowedBook.Date_borrowed)
		helper.PanicIfError(err)
		return BorrowedBook, nil
	} else {
		return BorrowedBook, err
	}
}

func (repository *BorrowedBookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.BorrowedBooksView {
	SQL := `SELECT id_borrowed, borrower.name as borrower_name, book.title as book_title, quantity_borrowed, quantity_back, date_borrowed FROM borrowed_books
		INNER JOIN borrower ON borrower.id = borrowed_books.id_borrower
		INNER JOIN book ON book.id = borrowed_books.id_book`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var BorrowedBooks []domain.BorrowedBooksView
	for rows.Next() {
		BorrowedBook := domain.BorrowedBooksView{}
		err := rows.Scan(&BorrowedBook.Id_borrowed,
			&BorrowedBook.Borrower_name,
			&BorrowedBook.Book_title,
			&BorrowedBook.Quantity_borrowed,
			&BorrowedBook.Quantity_back,
			&BorrowedBook.Date_borrowed)
		helper.PanicIfError(err)
		BorrowedBooks = append(BorrowedBooks, BorrowedBook)
	}
	return BorrowedBooks
}
