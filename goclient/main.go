package main

import
(
	"./dataTypes"
	"./tcpHttpClient"
)

func main() {
	// Make channels
	 tcpDataStreamCh := make(chan dataTypes.TcpDataStream)
	 tcpHttpClientStatusCh := make(chan dataTypes.TcpHttpClientStatus)
	 clientRequestCh := make(chan dataTypes.ClientRequest)

	 go tcpHttpClient.TcpHttpClientMain(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)
	 myReq := dataTypes.ClientRequest{Request:dataTypes.Start}
	 clientRequestCh<-myReq
	 select{}
}
