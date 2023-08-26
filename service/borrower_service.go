package service

import (
	"buku/model/web"
	"context"
)

type BorrowerService interface {
	Create(ctx context.Context, request web.BorrowerCreateRequest) web.BorrowerResponse
	Update(ctx context.Context, request web.BorrowerUpdateRequest) web.BorrowerResponse
	Delete(ctx context.Context, borrowerId int)
	FindById(ctx context.Context, borrowerId int) web.BorrowerResponse
	FindAll(ctx context.Context) []web.BorrowerResponse
}
