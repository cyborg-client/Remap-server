package tcphttpclient

import (
	"bytes"
	"encoding/json"
	"github.com/cyborg-client/Remap-server/config"
	"github.com/cyborg-client/Remap-server/datatypes"
	"log"
	"net/http"
)

// requestOptions is a struct representing the data which is sent to the MEA Server.
type requestOptions struct {
	SampleRate    int `json:"sample_rate""`
	SegmentLength int `json:"segment_length""`
}

// requestRemoteServer sends a Start / Stop post request to the MEA HTTP server.
func requestRemoteServer(start bool, sampleRate int, segmentLength int) bool {
	// Generate the URI based on start parameter
	var url string
	if start {
		url = "http://" + config.MEAServerAddress + ":" + config.MEAServerHTTPPort + "/start"
	} else {
		url = "http://" + config.MEAServerAddress + ":" + config.MEAServerHTTPPort + "/stop"
	}

	// Generate json
	reqOption := requestOptions{SampleRate: sampleRate, SegmentLength: segmentLength}
	b, err := json.Marshal(reqOption)
	if err != nil {
		log.Println("Error in http.go: ", err)
		panic(err)
	}

	var jsonStr = []byte(b)
	buf := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		log.Println("Error in http.go: ", err)
		panic(err)
	}
	defer resp.Body.Close() // Close HTTP when done

	// If HTTP response = 200 => success
	// if HTTP response != 200 => failure
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}
// httpMain defines the entrypoint for the http client, which sends START/STOP to the MEA server. Requires clientRequestCh, which it
// accepts Start or Stop, defining to send a start or stop request to the MEA server. Requires startStopTcpCh, which it uses to start/stop
// the tcp connection.
func httpMain(
	clientRequestCh <-chan datatypes.ClientRequest,
	startStopTcpCh chan<- startStopTcp,
) {
	for {
		select {
		case req := <-clientRequestCh:
			if req.Request == datatypes.Start {
				if requestRemoteServer(true, config.SampleRate, config.SegmentLength) {
					startStopTcpCh <- datatypes.Start
				}
			} else if req.Request == datatypes.Stop {
				requestRemoteServer(false, 0, 0)
			}
		}
	}

}
