package controller

import (
	"buku/helper"
	"buku/model/web"
	"buku/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BorrowerControllerImpl struct {
	BorrowerService service.BorrowerService
}

func NewBorrowerController(borrowerService service.BorrowerService) BorrowerController {
	return &BorrowerControllerImpl{BorrowerService: borrowerService}
}

func (controller *BorrowerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	borrowerCreateRequest := web.BorrowerCreateRequest{}
	helper.FromRequestToBody(request, &borrowerCreateRequest)

	borrowerResponse := controller.BorrowerService.Create(request.Context(), borrowerCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   borrowerResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	borrowerUpdateRequest := web.BorrowerUpdateRequest{}
	helper.FromRequestToBody(request, &borrowerUpdateRequest)

	borrowerId := params.ByName("borrowerId")
	id, err := strconv.Atoi(borrowerId)
	helper.PanicIfError(err)
	borrowerUpdateRequest.Id = id

	borrowerResponse := controller.BorrowerService.Update(request.Context(), borrowerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   borrowerResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	borrowerId := params.ByName("borrowerId")
	id, err := strconv.Atoi(borrowerId)
	helper.PanicIfError(err)

	controller.BorrowerService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Deleted",
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	borrowerId := params.ByName("borrowerId")
	id, err := strconv.Atoi(borrowerId)
	helper.PanicIfError(err)

	borrowerResponse := controller.BorrowerService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   borrowerResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *BorrowerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	borrowerResponse := controller.BorrowerService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   borrowerResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}
