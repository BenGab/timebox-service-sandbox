package main

import (
	"net/http"

	controller "github.com/bengab/timebox-service/src/controllers"
)

func main() {
	ctrl := controller.NewTimeController()

	http.Handle("/api/timestamp", ctrl)
	http.ListenAndServe(":8088", nil)
}
