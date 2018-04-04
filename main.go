package main

import (
	"fmt"
	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/datatypes"
	"github.com/cyborg-client/client/robotserver"
	"github.com/cyborg-client/client/tcphttpclient"
)

func main() {
	fmt.Println("Started application")

	// Make channels
	tcpDataStreamCh := make(chan tcphttpclient.TcpDataStream, 100)
	tcpHttpClientStatusCh := make(chan tcphttpclient.Status)
	clientRequestCh := make(chan datatypes.ClientRequest)

	go tcphttpclient.TcpHttpClient(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)

	myReq := datatypes.ClientRequest{Request: tcphttpclient.Start}
	myReqS := datatypes.ClientRequest{Request: tcphttpclient.Stop}
	clientRequestCh <- myReqS
	clientRequestCh <- myReq

	//run data parser
	timeStampChannel := make(chan []int64, 100)
	go analysis.Main(timeStampChannel, tcpDataStreamCh)
	go robotserver.Main(timeStampChannel)
	select {}
}
