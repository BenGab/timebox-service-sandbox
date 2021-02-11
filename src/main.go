package main

import (
	"net/http"
	"sync"

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

	wg.Wait()
}
