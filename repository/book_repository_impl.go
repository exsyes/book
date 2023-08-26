package repository

import (
	"buku/helper"
	"buku/model/domain"
	"context"
	"database/sql"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (Repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := `INSERT INTO book(Title, Quantity) VALUES ( ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, book.Title, book.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)

	return book
}

func (Repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := `UPDATE book SET Title = ? , Quantity = ? WHERE Id = ?`
	_, err := tx.ExecContext(ctx, SQL, book.Title, book.Quantity, book.Id)
	helper.PanicIfError(err)

	return book
}

func (Repository *BookRepositoryImpl) UpdateQty(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := `UPDATE book SET Quantity = ? WHERE Id = ?`
	_, err := tx.ExecContext(ctx, SQL, book.Quantity, book.Id)
	helper.PanicIfError(err)

	return book
}

func (Repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	SQL := `DELETE FROM book WHERE Id = ?`
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	helper.PanicIfError(err)

}

func (Repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	SQL := `SELECT Id ,Title, Quantity FROM book WHERE Id = ?`
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	book := domain.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Quantity)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, err
	}
}

func (Repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := `SELECT * FROM book`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Quantity)
		helper.PanicIfError(err)
		books = append(books, book)
	}

	return books
}
