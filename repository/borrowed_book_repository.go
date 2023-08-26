package repository

import (
	"buku/model/domain"
	"context"
	"database/sql"
)

type BorrowedBookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, BorrowedBook domain.BorrowedBooks) domain.BorrowedBooks
	Update(ctx context.Context, tx *sql.Tx, BorrowedBook domain.BorrowedBooks) domain.BorrowedBooks
	Delete(ctx context.Context, tx *sql.Tx, BorrowedBook domain.BorrowedBooks)
	FindById(ctx context.Context, tx *sql.Tx, BorrowedBookId int) (domain.BorrowedBooks, error)
	FindByIdView(ctx context.Context, tx *sql.Tx, BorrowedBookId int) (domain.BorrowedBooksView, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BorrowedBooksView
}
