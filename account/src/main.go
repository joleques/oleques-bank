package main

import (
	"github.com/joleques/oleques-bank/account/src/application/api"
	"github.com/joleques/oleques-bank/account/src/infrastructure/log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	status := "API Account Up - Version 01.00.00"
	log.Warning(status, time.Now())
	wg.Add(1)
	go api.Start(status)
	wg.Wait()
}
