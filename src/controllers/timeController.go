package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bengab/timebox-service/src/logic"
)

type timeController struct {
	service logic.TimestampService
}

func NewTimeController() timeController {
	return timeController{service: logic.NewtimestampService()}
}

func (tc timeController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid content type"))
		return
	}

	switch req.Method {
	case "POST":
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal("invalid body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		str := string(b)

		fmt.Println("BODY")
		fmt.Println(str)

		if err != nil {
			log.Fatal("invalid body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tc.service.SetTimestamp(str)
		w.WriteHeader(http.StatusNoContent)

	case "GET":
		timestr := tc.service.GetTimeStamp()

		if timestr == nil {
			w.Write([]byte(time.Now().String()))
		} else {
			w.Write([]byte(*timestr))
		}
	}
}
