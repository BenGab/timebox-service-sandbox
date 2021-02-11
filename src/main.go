package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/bengab/timebox-service/src/client"
	controller "github.com/bengab/timebox-service/src/controllers"
)

func startWebService(wg *sync.WaitGroup) {
	ctrl := controller.NewTimeController()
	defer wg.Done()
	http.Handle("/api/timestamp", ctrl)
	http.ListenAndServe(":8088", nil)

}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go startWebService(wg)

	store := client.NewStore("http://localhost:8088/api/timestamp")
	stopChan := make(chan struct{})
	go store.Start(stopChan)

	store.WriteChan <- time.Now()
	go store.ReadTimeFromServer()

	wg.Wait()
	stopChan <- struct{}{}
}
