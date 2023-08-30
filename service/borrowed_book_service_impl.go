package service

import (
	"buku/exception"
	"buku/helper"
	"buku/model/domain"
	"buku/model/web"
	"buku/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type BorrowedBookServiceImpl struct {
	Db                     *sql.DB
	BorrowedBookrepository repository.BorrowedBookRepository
	BookRepository         repository.BookRepository
	Validate               *validator.Validate
}

func NewBorrowedBookService(db *sql.DB, borrowedBookrepository repository.BorrowedBookRepository, bookRepository repository.BookRepository, validate *validator.Validate) BorrowedBookService {
	return &BorrowedBookServiceImpl{
		Db:                     db,
		BorrowedBookrepository: borrowedBookrepository,
		BookRepository:         bookRepository,
		Validate:               validate,
	}
}

func (service *BorrowedBookServiceImpl) Create(ctx context.Context, request web.BorrowedBookCreateRequest) web.BorrowedBookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id_book)
	helper.PanicIfError(err)
	book.Quantity = book.Quantity - request.Quantity_borrowed
	book = service.BookRepository.UpdateQty(ctx, tx, book)

	BorrowedBook := domain.BorrowedBooks{
		Id_book:           request.Id_book,
		Id_borrower:       request.Id_borrower,
		Quantity_borrowed: request.Quantity_borrowed,
	}
	BorrowedBook = service.BorrowedBookrepository.Save(ctx, tx, BorrowedBook)
	return helper.ToBorrowedBooksResponse(BorrowedBook)
}

func (service *BorrowedBookServiceImpl) Update(ctx context.Context, request web.BorrowedBookUpdateRequest) web.BorrowedBookViewResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id_book)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	book.Quantity = book.Quantity + request.Quantity_back
	book = service.BookRepository.UpdateQty(ctx, tx, book)

	BorrowedBook, err := service.BorrowedBookrepository.FindByIdView(ctx, tx, request.Id_borrowed)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	BorrowedBook.Quantity_borrowed = BorrowedBook.Quantity_borrowed - request.Quantity_back
	BorrowedBook.Quantity_back = BorrowedBook.Quantity_back + request.Quantity_back
	BorrowedBook = service.BorrowedBookrepository.Update(ctx, tx, BorrowedBook)

	return helper.ToBorrowedBooksViewResponse(BorrowedBook)
}

func (service *BorrowedBookServiceImpl) Delete(ctx context.Context, borrowedId int) {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	BorrowedBook, err := service.BorrowedBookrepository.FindById(ctx, tx, borrowedId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BorrowedBookrepository.Delete(ctx, tx, BorrowedBook)
}

func (service *BorrowedBookServiceImpl) FindByIdView(ctx context.Context, borrowedId int) web.BorrowedBookViewResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	BorrowedBook, err := service.BorrowedBookrepository.FindByIdView(ctx, tx, borrowedId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBorrowedBooksViewResponse(BorrowedBook)

}

func (service *BorrowedBookServiceImpl) FindById(ctx context.Context, borrowedId int) web.BorrowedBookResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	BorrowedBook, err := service.BorrowedBookrepository.FindById(ctx, tx, borrowedId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBorrowedBooksResponse(BorrowedBook)

}

func (service *BorrowedBookServiceImpl) FindAll(ctx context.Context) []web.BorrowedBookViewResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	BorrowedBooks := service.BorrowedBookrepository.FindAll(ctx, tx)
	return helper.ToBorrowedBooksViewResponses(BorrowedBooks)
}
