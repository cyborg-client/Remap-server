package main

import
(
	"github.com/vegardbotnen/EiT-Client/datatypes"
	"github.com/vegardbotnen/EiT-Client/tcphttpclient"
)

func main() {
	// Make channels
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
}
