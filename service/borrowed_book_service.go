package service

import (
	"buku/model/web"
	"context"
)

type BorrowedBookService interface {
	Create(ctx context.Context, request web.BorrowedBookCreateRequest) web.BorrowedBookResponse
	Update(ctx context.Context, request web.BorrowedBookUpdateRequest) web.BorrowedBookResponse
	Delete(ctx context.Context, borrowedId int)
	FindByIdView(ctx context.Context, borrowedId int) web.BorrowedBookViewResponse
	FindAll(ctx context.Context) []web.BorrowedBookViewResponse
}
