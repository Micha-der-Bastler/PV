package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Router (links all HTTP-Requests with special url pattern to a special Request-Handler)
	http.HandleFunc("/on", requestHandlerOn)
	http.HandleFunc("/off", requestHandlerOff)

	// Acceptance of HTTP-Requests
	http.ListenAndServe(":80", nil)
}

// Request-Handler
func requestHandlerOn(resp http.ResponseWriter, req *http.Request) {
	resp2, err := http.Get("http://shelly1-76f2fa/relay/0?turn=on")
	if err != nil {
		log.Fatalln(err)
	}
	// Hi Angiiii
	defer resp2.Body.Close()

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(string(body))
	resp.Write([]byte(body))
}
func requestHandlerOff(resp http.ResponseWriter, req *http.Request) {
	resp2, err := http.Get("http://shelly1-76f2fa/relay/0?turn=off")
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
