package tcphttpclient

import (
	"github.com/cyborg-client/client/datatypes"
	"github.com/cyborg-client/client/config"
	"fmt"
	"net/http"
	"log"
	"bytes"
	"encoding/json"
)

// requestOptions Struct representing the data which is sent to the MEA Server.
type requestOptions struct {
	SampleRate int `json:"sample_rate""`
	SegmentLength int `json:"segment_length""`
}

func requestRemoteServer(start bool, sampleRate int, segmentLength int) (bool) {
	var url string
	if start {
		url = "http://" + config.MEAServerAddress + "/start"
	}else {
		url = "http://" + config.MEAServerAddress + "/stop"
	}

	// Generate json
	reqOption := requestOptions{SampleRate:sampleRate, SegmentLength:segmentLength}
	b, err := json.Marshal(reqOption)
	if err != nil {
		// TODO: Handle error
		fmt.Println("error:", err)
	}

	var jsonStr = []byte(b)
	buf := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		// Error
		// TODO : Handle error
		log.Println("Error in http.go: ", err)
		return false
	}
	defer resp.Body.Close() // Close HTTP when done

	// DEBUG
	// TODO: Remove debug when the communication with the server is functioning
	bufT := new(bytes.Buffer)
	bufT.ReadFrom(resp.Body)
	newStr := bufT.String()
	fmt.Println(newStr)
	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func httpMain(
	statusTcpCh <-chan statusTcp,
	clientRequestCh <-chan datatypes.ClientRequest,
	startStopTcpCh chan<- startStopTcp,
	tcpHttpClientStatusCh chan<- datatypes.TcpHttpClientStatus,
) {
	for{
		select {
		case req := <-clientRequestCh:
			if req.Request == datatypes.Start {
				fmt.Println("start")
				if requestRemoteServer(true, req.Options["sample_rate"], req.Options["segment_length"]){
				}

			} else if req.Request == datatypes.Stop {
				fmt.Println("Stop")
				requestRemoteServer(false, 0, 0)
			}
		}
	}

}
