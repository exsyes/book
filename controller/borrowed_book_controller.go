package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BorrowedBookController interface {
	Create(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(response http.ResponseWriter, request *http.Request, params httprouter.Params)
}
