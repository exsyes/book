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

type BorrowerServiceImpl struct {
	Db                  *sql.DB
	BorrowerRespository repository.BorrowerRepository
	Validate            *validator.Validate
}

func NewBorrowerService(db *sql.DB, borrowerRespository repository.BorrowerRepository, validate *validator.Validate) BorrowerService {
	return &BorrowerServiceImpl{
		Db:                  db,
		BorrowerRespository: borrowerRespository,
		Validate:            validate,
	}
}

func (service *BorrowerServiceImpl) Create(ctx context.Context, request web.BorrowerCreateRequest) web.BorrowerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	borrower := domain.Borrower{
		Name: request.Name,
	}

	borrower = service.BorrowerRespository.Save(ctx, tx, borrower)
	return helper.ToBorrowerResponse(borrower)
}

func (service *BorrowerServiceImpl) Update(ctx context.Context, request web.BorrowerUpdateRequest) web.BorrowerResponse {
	err := service.Validate.Struct(request)
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	borrower, err := service.BorrowerRespository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	borrower.Name = request.Name

	borrower = service.BorrowerRespository.Update(ctx, tx, borrower)

	return helper.ToBorrowerResponse(borrower)
}

func (service *BorrowerServiceImpl) Delete(ctx context.Context, borrowerId int) {

	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	borrower, err := service.BorrowerRespository.FindById(ctx, tx, borrowerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BorrowerRespository.Delete(ctx, tx, borrower)
}

func (service *BorrowerServiceImpl) FindById(ctx context.Context, borrowerId int) web.BorrowerResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	borrower, err := service.BorrowerRespository.FindById(ctx, tx, borrowerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBorrowerResponse(borrower)
}

func (service *BorrowerServiceImpl) FindAll(ctx context.Context) []web.BorrowerResponse {
	tx, err := service.Db.Begin()
	helper.PanicIfError(err)

	borrowers := service.BorrowerRespository.FindAll(ctx, tx)

	return helper.ToBorrowerResponses(borrowers)
}
