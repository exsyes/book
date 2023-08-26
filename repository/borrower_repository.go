package repository

import (
	"buku/model/domain"
	"context"
	"database/sql"
)

type BorrowerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, borrower domain.Borrower) domain.Borrower
	Update(ctx context.Context, tx *sql.Tx, borrower domain.Borrower) domain.Borrower
	Delete(ctx context.Context, tx *sql.Tx, borrower domain.Borrower)
	FindById(ctx context.Context, tx *sql.Tx, borrowerId int) (domain.Borrower, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Borrower
}
