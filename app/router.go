package app

import (
	"buku/controller"
	"buku/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController,
	borrowerController controller.BorrowerController,
	borrowedBookController controller.BorrowedBookController) *httprouter.Router {

	router := httprouter.New()
	//book
	router.GET("/book", bookController.FindAll)
	router.GET("/book/:bookId", bookController.FindById)
	router.POST("/book", bookController.Create)
	router.PUT("/book/:bookId", bookController.Update)
	router.PUT("/book/:bookId/quantity", bookController.UpdateQty)
	router.DELETE("/book/:bookId", bookController.Delete)

	//borrower
	router.GET("/borrower", borrowerController.FindAll)
	router.GET("/borrower/:borrowerId", borrowerController.FindById)
	router.POST("/borrower", borrowerController.Create)
	router.PUT("/borrower/:borrowerId", borrowerController.Update)
	router.DELETE("/borrower/:borrowerId", borrowerController.Delete)

	//borrowed
	router.GET("/borrowed", borrowedBookController.FindAll)
	router.GET("/borrowed/:id_borrowed", borrowedBookController.FindById)
	router.POST("/borrowed", borrowedBookController.Create)
	router.PUT("/borrowed/:id_borrowed", borrowedBookController.Update)
	router.DELETE("/borrowed/:id_borrowed", borrowedBookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
