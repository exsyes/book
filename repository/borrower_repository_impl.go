package repository

import (
	"buku/helper"
	"buku/model/domain"
	"context"
	"database/sql"
)

type BorrowerRepositoryImpl struct {
}

func NewBorrowerRepository() BorrowerRepository {
	return &BorrowerRepositoryImpl{}
}

func (Repository *BorrowerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, borrower domain.Borrower) domain.Borrower {
	SQL := `INSERT INTO borrower(name) VALUES (?)`
	result, err := tx.ExecContext(ctx, SQL, borrower.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	borrower.Id = int(id)
	return borrower
}

func (Repository *BorrowerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, borrower domain.Borrower) domain.Borrower {
	SQL := `UPDATE borrower SET Name = ? WHERE Id = ?`
	_, err := tx.ExecContext(ctx, SQL, borrower)
	helper.PanicIfError(err)

	return borrower
}

func (Repository *BorrowerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, borrower domain.Borrower) {
	SQL := `DELETE FROM borrower WHERE Id = ?`
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}

func (Repository *BorrowerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, borrowerId int) (domain.Borrower, error) {
	SQL := `SELECT Id ,Name FROM borrower WHERE Id = ?`
	rows, err := tx.QueryContext(ctx, SQL, borrowerId)
	helper.PanicIfError(err)
	defer rows.Close()

	var borrower domain.Borrower
	if rows.Next() {
		err := rows.Scan(&borrower.Id, &borrower.Name)
		helper.PanicIfError(err)
		return borrower, nil
	} else {
		return borrower, err
	}
}

func (Repository *BorrowerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Borrower {
	SQL := `SELECT Id ,Name FROM borrower`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var borrowers []domain.Borrower
	for rows.Next() {
		var borrower domain.Borrower
		rows.Scan(&borrower.Id, &borrower.Name)
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}
