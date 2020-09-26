package main

import (
	"PV/power/powerRepository/powerRepositoryRest"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	pRR := powerRepositoryRest.NewPowerRepositoryRest()
	logrus.Info(pRR.GetPower("http://shelly1pm-BA0F5F"))

	// Router (links all HTTP-Requests with special url pattern to a special Request-Handler)
	http.HandleFunc("/on", requestHandlerOn)
	http.HandleFunc("/off", requestHandlerOff)
	http.HandleFunc("/version", requestHandlerVersion)

	// Acceptance of HTTP-Requests
	http.ListenAndServe(":80", nil)
}

// Request-Handler
func requestHandlerOn(resp http.ResponseWriter, req *http.Request) {
	resp2, err := http.Get("http://shelly1pm-BA0F5F/relay/0?turn=on")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp2.Body.Close()

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(string(body))
	resp.Write([]byte(body))
}
func requestHandlerOff(resp http.ResponseWriter, req *http.Request) {
	resp2, err := http.Get("http://shelly1pm-BA0F5F/relay/0?turn=off")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp2.Body.Close()

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(string(body))
	resp.Write([]byte(body))
}
func requestHandlerVersion(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "11.04.2020 20:34")
}
