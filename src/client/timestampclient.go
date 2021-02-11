package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type store struct {
	readChan  chan string
	WriteChan chan time.Time
	serverURL string
}

func (s *store) Start(stopChan <-chan struct{}) {
	for {
		select {
		case result, ok := <-s.readChan:
			if !ok {
				break
			}
			fmt.Println("From channel process")
			fmt.Println(result)

		case timestamp, ok := <-s.WriteChan:
			if !ok {
				break
			}
			s.writeTimeToServer(&timestamp)

		case <-stopChan:
			fmt.Println("STOP")
			close(s.readChan)
			close(s.WriteChan)
			break
		}

	}
}

func (s *store) writeTimeToServer(time *time.Time) {
	_, err := http.Post(s.serverURL, "text/plain", bytes.NewBufferString(time.String()))

	if err != nil {
		log.Fatalln(err)
	}
}

func (s *store) ReadTimeFromServer() {
	req, err := http.NewRequest("GET", s.serverURL, &bytes.Buffer{})

	if err != nil {
		log.Fatalln(err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
		return
	}

	s.readChan <- string(b)
}

func NewStore(aserverURL string) *store {
	return &store{
		readChan:  make(chan string),
		WriteChan: make(chan time.Time),
		serverURL: aserverURL,
	}
}
