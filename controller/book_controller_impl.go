package controller

import (
	"buku/helper"
	"buku/model/web"
	"buku/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookservice service.BookService) BookController {
	return &BookControllerImpl{BookService: bookservice}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookCreateRequest := web.BookCreateRequest{}
	helper.FromRequestToBody(request, &bookCreateRequest)

	bookResponse := controller.BookService.Create(request.Context(), bookCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   bookResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}
	helper.FromRequestToBody(request, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	bookResponse := controller.BookService.Update(request.Context(), bookUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   bookResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BookControllerImpl) UpdateQty(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookUpdateQtyRequest := web.BookUpdateQtyRequest{}
	helper.FromRequestToBody(request, &bookUpdateQtyRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateQtyRequest.Id = id

	bookResponse := controller.BookService.UpdateQty(request.Context(), bookUpdateQtyRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   bookResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success delete",
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Succcess",
		Data:   bookResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookResponses := controller.BookService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   bookResponses,
	}

	helper.WriteToResponse(writer, webResponse)
}
