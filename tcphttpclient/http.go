package tcphttpclient

import (
	"github.com/vegardbotnen/EiT-Client/datatypes"
	"github.com/vegardbotnen/EiT-Client/config"
	"fmt"
	"net/http"
	"log"
	"bytes"
)

func requestRemoteServer(start bool) (bool) {
	var url string
	if start {
		url = "http://" + config.MEAServerAddress + "/start"
	}else {
		url = "http://" + config.MEAServerAddress + "/stop"
	}
	var jsonStr = []byte(`{"sample_rate":1000, "segment_length":10}`)
	buf := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		// Error
		log.Println("Error in http.go: ", err)
		return false
	}
	defer resp.Body.Close() // Close HTTP when done

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
				if requestRemoteServer(true){
				}

			} else if req.Request == datatypes.Stop {
				fmt.Println("Stop")
				requestRemoteServer(false)
			}
		}
	}

}
