package websocketserver

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/cyborg-client/client/config"
	"log"
)

type stimulateRequest struct {
	Frequency int `json:"frequency""`
	Duration int `json:"duration""`
	Channel int `json:"channel""`
}

// https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
// stimulateServer is the handler for HTTP POST requests to stimulate the MEA server. Verifies that the incoming POST request
// is valid, and the json payload is valid, before sending it to the MEA server.
func stimulateServer(rw http.ResponseWriter, req *http.Request){
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var t stimulateRequest
	err = json.Unmarshal(body, &t)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// We are just redirecting the data. Proceeding by sending it. We have verified that the JSON is valid
	var jsonStr = []byte(body)
	buf := bytes.NewBuffer(jsonStr)
	url := "http://" + config.MEAServerAddress + ":" + config.MEAServerHTTPPort + "/stimulate"
	resp, err := http.Post(url, "application/json", buf)
	defer resp.Body.Close()
	if err != nil {
		log.Println("Error in http.go: ", err)
		panic(err)
	}
	log.Println(resp.StatusCode)
}
