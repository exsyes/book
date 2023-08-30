package service

import (
	"buku/model/web"
	"context"
)

type BorrowedBookService interface {
	Create(ctx context.Context, request web.BorrowedBookCreateRequest) web.BorrowedBookResponse
	Update(ctx context.Context, request web.BorrowedBookUpdateRequest) web.BorrowedBookViewResponse
	Delete(ctx context.Context, borrowedId int)
	FindByIdView(ctx context.Context, borrowedId int) web.BorrowedBookViewResponse
	FindById(ctx context.Context, borrowedId int) web.BorrowedBookResponse
	FindAll(ctx context.Context) []web.BorrowedBookViewResponse
}
