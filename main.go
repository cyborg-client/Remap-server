package main

import (
	"fmt"
	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/datatypes"
	"github.com/cyborg-client/client/websocketserver"
	"github.com/cyborg-client/client/tcphttpclient"
)

func main() {
	fmt.Println("Started application")

	// Make channels
	tcpDataStreamCh := make(chan tcphttpclient.Segment, 100)
	clientRequestCh := make(chan datatypes.ClientRequest)

	go tcphttpclient.Main(tcpDataStreamCh, clientRequestCh)

	myReq := datatypes.ClientRequest{Request: datatypes.Start}
	myReqS := datatypes.ClientRequest{Request: datatypes.Stop}
	clientRequestCh <- myReqS
	clientRequestCh <- myReq

	//run data parser
	timeStampChannel := make(chan []int64, 100)
	go analysis.Main(timeStampChannel, tcpDataStreamCh)
	go websocketserver.Main(timeStampChannel)
	select {}
}
