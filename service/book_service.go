package service

import (
	"buku/model/web"
	"context"
)

type BookService interface {
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	UpdateQty(ctx context.Context, request web.BookUpdateQtyRequest) web.BookResponse
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) web.BookResponse
	FindAll(ctx context.Context) []web.BookResponse
}
