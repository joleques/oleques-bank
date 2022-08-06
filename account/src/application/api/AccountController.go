package api

import (
	"github.com/go-chi/chi"
	renderChi "github.com/go-chi/render"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	"github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/application/userCase"
	"github.com/joleques/oleques-bank/account/src/infrastructure/log"
	"github.com/joleques/oleques-bank/account/src/infrastructure/login"
	"github.com/joleques/oleques-bank/account/src/infrastructure/repository"
	"net/http"
	"time"
)

func Get(writer http.ResponseWriter, request *http.Request) {
	initProcess := time.Now()
	userCase, err := userCase.NewGetAccount(&converters.AccountConverter{}, repository.AccountMemory{})
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	account, err := userCase.Get(chi.URLParam(request, "account_id"))
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	render.JSON(writer, 200, dto.BalanceDTO{ID: account.ID, Balance: account.Balance})
}

func Update(writer http.ResponseWriter, request *http.Request) {
	initProcess := time.Now()
	userCase, err := userCase.NewUpdateBalance(repository.AccountMemory{})
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	data := &dto.BalanceDTO{}
	if err := renderChi.Bind(request, data); err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 400, initProcess)
		return
	}
	data.ID = chi.URLParam(request, "account_id")
	err = userCase.Update(*data)
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	render.JSON(writer, 200, dto.ApiResponse{StatusCode: 200, Message: "Update success"})
}

func List(writer http.ResponseWriter, request *http.Request) {
	initProcess := time.Now()
	userCase, err := userCase.NewListAccount(&converters.AccountConverter{}, repository.AccountMemory{})
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	list := userCase.List()
	render.JSON(writer, 200, list)
}

func CreateAccount(writer http.ResponseWriter, request *http.Request) {
	initProcess := time.Now()
	data := &dto.AccountDTO{}
	if err := renderChi.Bind(request, data); err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 400, initProcess)
		return
	}
	useCase, err := userCase.NewCreateAccount(&converters.AccountConverter{}, repository.AccountMemory{}, login.LoginService{})
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	err = useCase.Create(*data)
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	result := dto.ApiResponse{StatusCode: 201, Message: "Account created successfully"}
	render.JSON(writer, 201, result)
}

func error(writer http.ResponseWriter, messageError string, status int, initProcess time.Time) {
	response := dto.ApiResponse{StatusCode: 400, Message: messageError}
	render.JSON(writer, status, response)
	log.PrintLog(dto.AccountDTO{}, response, initProcess)
}
