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

type BookServiceImpl struct {
	Db             *sql.DB
	BookRepository repository.BookRepository
	Validate       *validator.Validate
}

func NewBookService(db *sql.DB, bookRepository repository.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		Db:             db,
		BookRepository: bookRepository,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := domain.Book{
		Title:    request.Title,
		Quantity: request.Quantity,
	}

	book = service.BookRepository.Save(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	book.Title = request.Title
	book.Quantity = request.Quantity

	book = service.BookRepository.Update(ctx, tx, book)
	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) UpdateQty(ctx context.Context, request web.BookUpdateQtyRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	book.Quantity = request.Quantity

	book = service.BookRepository.UpdateQty(ctx, tx, book)
	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.BookRepository.Delete(ctx, tx, book)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) web.BookResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAll(ctx, tx)
	return helper.ToBookResponses(books)
}
