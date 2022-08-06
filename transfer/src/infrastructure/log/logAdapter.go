package log

import (
	"encoding/json"
	"fmt"
	"github.com/joleques/oleques-bank/transfer/src/application/dto"
	"time"
)

type Header struct {
	ServiceName    string `json:"serviceName"`
	Status         string `json:"status"`
	StatusCode     int    `json:"statusCode"`
	ProcessingTime int64  `json:"processingTime"`
}

type Log struct {
	Ts      string `json:"ts"`
	Header  Header `json:"header"`
	Message string `json:"message"`
}

const SERVICE_NAME = "transfer-service"

func PrintLog(request dto.TransferDTO, response dto.ApiResponse, initProcess time.Time) {
	status := "SUCCESS"
	if response.StatusCode > 300 {
		status = "ERROR"
	}

	header := Header{ServiceName: SERVICE_NAME, StatusCode: response.StatusCode, Status: status, ProcessingTime: time.Since(initProcess).Milliseconds()}
	log := Log{Header: header, Message: response.Message, Ts: time.Now().UTC().Format(time.RFC3339)}
	print(log)
}

func Warning(msg string, initProcess time.Time) {
	header := Header{ServiceName: SERVICE_NAME, Status: "WARNING", ProcessingTime: time.Since(initProcess).Milliseconds()}
	log := Log{Header: header, Message: msg, Ts: time.Now().UTC().Format(time.RFC3339)}
	print(log)
}

func print(log Log) {
	jsonByte, _ := json.Marshal(log)
	fmt.Println(string(jsonByte))
}
