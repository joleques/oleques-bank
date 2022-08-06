package main

import (
	"github.com/joleques/oleques-bank/transfer/src/application/api"
	"github.com/joleques/oleques-bank/transfer/src/infrastructure/log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	status := "API Transfer Up - Version 01.00.00"
	log.Warning(status, time.Now())
	wg.Add(1)
	go api.Start(status)
	wg.Wait()
}
