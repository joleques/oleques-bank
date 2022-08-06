package api

import (
	renderChi "github.com/go-chi/render"
	"github.com/joleques/oleques-bank/transfer/src/application/dto"
	useCase2 "github.com/joleques/oleques-bank/transfer/src/application/useCase"
	"github.com/joleques/oleques-bank/transfer/src/infrastructure/accountAdapter"
	"github.com/joleques/oleques-bank/transfer/src/infrastructure/log"
	repository2 "github.com/joleques/oleques-bank/transfer/src/infrastructure/repository"
	"net/http"
	"time"
)

func Transfer(writer http.ResponseWriter, request *http.Request) {
	initProcess := time.Now()
	data := &dto.TransferDTO{}
	if err := renderChi.Bind(request, data); err != nil {
		error(writer, "Invalid request.... Reason: "+err.Error(), 400, initProcess)
		return
	}
	service := accountAdapter.AccountAdapter{}
	repository := repository2.TransferMemory{}
	useCase, err := useCase2.NewTransferBetweenAccounts(service, repository)
	if err != nil {
		error(writer, "Invalid request. Reason: "+err.Error(), 500, initProcess)
		return
	}
	err = useCase.Transfer(*data)
	if err != nil {
		error(writer, "Business error. Reason: "+err.Error(), 500, initProcess)
		return
	}
	result := dto.ApiResponse{StatusCode: 201, Message: "transfer performed successfully"}
	render.JSON(writer, 201, result)
}

func error(writer http.ResponseWriter, messageError string, status int, initProcess time.Time) {
	response := dto.ApiResponse{StatusCode: status, Message: messageError}
	render.JSON(writer, status, response)
	log.PrintLog(dto.TransferDTO{}, response, initProcess)
}
