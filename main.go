package main

import (
	"buku/app"
	"buku/controller"
	"buku/helper"
	"buku/repository"
	"buku/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDb()
	validate := validator.New()

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(db, bookRepository, validate)
	bookController := controller.NewBookController(bookService)

	borrowerRepository := repository.NewBorrowerRepository()
	borrowerService := service.NewBorrowerService(db, borrowerRepository, validate)
	borrowerController := controller.NewBorrowerController(borrowerService)

	borrowedBookRepository := repository.NewBorrowedBookRepository()
	borrowedBookService := service.NewBorrowedBookService(db, borrowedBookRepository, bookRepository, validate)
	borrowedBookController := controller.NewBorrowedBookControllerImpl(borrowedBookService)
	router := app.NewRouter(bookController, borrowerController, borrowedBookController)

	server := http.Server{
		Addr:    "localhost:9191",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
