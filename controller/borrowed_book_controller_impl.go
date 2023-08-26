package controller

import (
	"buku/helper"
	"buku/model/web"
	"buku/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BorrowedBookControllerImpl struct {
	service service.BorrowedBookService
}

func NewBorrowedBookControllerImpl(service service.BorrowedBookService) BorrowedBookController {
	return &BorrowedBookControllerImpl{service: service}
}

func (controller *BorrowedBookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	borrowedBookCreateRequest := web.BorrowedBookCreateRequest{}
	helper.FromRequestToBody(request, &borrowedBookCreateRequest)

	borrowedBookResponse := controller.service.Create(request.Context(), borrowedBookCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   borrowedBookResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowedBookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	borrowedBookUpdateRequest := web.BorrowedBookUpdateRequest{}
	helper.FromRequestToBody(request, &borrowedBookUpdateRequest)

	idBorrowed := params.ByName("id_borrowed")
	id, err := strconv.Atoi(idBorrowed)
	helper.PanicIfError(err)
	borrowedBookUpdateRequest.Id_borrowed = id
	borrowedBookResponse := controller.service.Update(request.Context(), borrowedBookUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   borrowedBookResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowedBookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idBorrowed := params.ByName("id_borrowed")
	id, err := strconv.Atoi(idBorrowed)
	helper.PanicIfError(err)
	controller.service.Delete(request.Context(), int(id))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowedBookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idBorrowed := params.ByName("id_borrowed")
	id, err := strconv.Atoi(idBorrowed)
	helper.PanicIfError(err)
	borrowedBookResponse := controller.service.FindByIdView(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   borrowedBookResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowedBookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	borrowedBookResponse := controller.service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   borrowedBookResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}
