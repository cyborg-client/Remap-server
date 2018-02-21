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

	 testConfig := make(map[string]int)
	 testConfig["sample_rate"] = 1000
	 testConfig["segment_length"] = 10

	 myReq := dataTypes.ClientRequest{Request:dataTypes.Start, Options:testConfig}
	 clientRequestCh <- myReq
	 /*
	 myReq := dataTypes.ClientRequest{Request:dataTypes.Start}
	 myReqStop := dataTypes.ClientRequest{Request:dataTypes.Stop}
	 clientRequestCh<-myReq
	 clientRequestCh<-myReqStop
	 clientRequestCh<-myReq
	 clientRequestCh<-myReqStop
*/
	 select{}
}
