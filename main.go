package main

import (
	"github.com/cyborg-client/client/goclient/config"
	"github.com/cyborg-client/client/goclient/robotserver"
)

func main() {
	// // Make channels
	// tcpDataStreamCh := make(chan dataTypes.TcpDataStream)
	// tcpHttpClientStatusCh := make(chan dataTypes.TcpHttpClientStatus)
	// clientRequestCh := make(chan dataTypes.ClientRequest)

	// go tcpHttpClient.TcpHttpClientMain(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)
	// myReq := dataTypes.ClientRequest{Request: dataTypes.Start}
	// myReqStop := dataTypes.ClientRequest{Request: dataTypes.Stop}
	// clientRequestCh <- myReq
	// clientRequestCh <- myReqStop
	// clientRequestCh <- myReq
	// clientRequestCh <- myReqStop
	// select {}

	// Robot Server
	robotserver.Main(config.RobotServerPort)
}
