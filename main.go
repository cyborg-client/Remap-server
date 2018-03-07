package main

import (
	"time"

	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/config"
	"github.com/cyborg-client/client/robotserver"
)

// "github.com/cyborg-client/client/datatypes"
// "github.com/cyborg-client/client/tcphttpclient"

func main() {
	// Make channels
	/*
	 tcpDataStreamCh := make(chan datatypes.TcpDataStream)
	 tcpHttpClientStatusCh := make(chan datatypes.TcpHttpClientStatus)
	 clientRequestCh := make(chan datatypes.ClientRequest)

	 go tcphttpclient.TcpHttpClientMain(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)
	 myReq := datatypes.ClientRequest{Request:datatypes.Start}
	 myReqStop := datatypes.ClientRequest{Request:datatypes.Stop}
	 clientRequestCh<-myReq
	 clientRequestCh<-myReqStop
	 clientRequestCh<-myReq
	 clientRequestCh<-myReqStop
	 select{}
	*/
	//run data parser
	timeStampChannel := make(chan []byte, 10)
	go analysis.Main(timeStampChannel)
	go robotserver.Main(config.RobotServerPort, timeStampChannel)
	time.Sleep(20000 * time.Millisecond)
}
