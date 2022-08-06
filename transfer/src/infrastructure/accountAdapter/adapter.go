package accountAdapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joleques/oleques-bank/transfer/src/domain"
	"io/ioutil"
	"net/http"
)

type AccountAdapter struct {
}

type ResponseAPI struct {
	StatusCode int64
	Message    string
}

type ResponseBalanceDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type RequestBalance struct {
	Balance float64 `json:"balance"`
}

func (adapter AccountAdapter) UpdateAccount(account domain.Account) error {
	url := "http://localhost:3000/accounts/" + account.Id + "/balance"
	dto := RequestBalance{Balance: account.Balance()}
	result, err := json.Marshal(dto)
	if err != nil {
		msg := fmt.Sprintf("error json convert. Reason %v", err.Error())
		return errors.New(msg)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(result))
	if err != nil {
		return err
	}
	request.Header.Add("content-type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bodyInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		responseAPI := &ResponseAPI{}
		err = json.Unmarshal(bodyInBytes, responseAPI)
		if err != nil {
			return err
		}
		return errors.New(responseAPI.Message)
	}
	return nil
}

func (adapter AccountAdapter) GetBalance(id string) (float64, error) {
	url := "http://localhost:3000/accounts/" + id + "/balance"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, err
	}
	request.Header.Add("content-type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0.0, err
	}
	defer response.Body.Close()
	bodyInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	if response.StatusCode == 404 {
		return 0.0, errors.New(string(bodyInBytes) + " Url: " + url)
	}
	if response.StatusCode != 200 {
		responseAPI := &ResponseAPI{}
		err = json.Unmarshal(bodyInBytes, responseAPI)
		if err != nil {
			return 0.0, errors.New("Unmarshal body - " + err.Error())
		}
		return 0.0, errors.New(responseAPI.Message)
	}
	responseAPI := &ResponseBalanceDTO{}
	err = json.Unmarshal(bodyInBytes, responseAPI)
	if err != nil {
		return 0.0, errors.New("Unmarshal body - " + err.Error())
	}
	return responseAPI.Balance, nil
}
