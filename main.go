package main

import
(
	"github.com/cyborg-client/client/datatypes"
	"github.com/cyborg-client/client/tcphttpclient"
)

func main() {
	// Make channels
	 tcpDataStreamCh := make(chan tcphttpclient.TcpDataStream)
	 tcpHttpClientStatusCh := make(chan tcphttpclient.TcpHttpClientStatus)
	 clientRequestCh := make(chan datatypes.ClientRequest)

	 go tcphttpclient.TcpHttpClientMain(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)

	 //myReq := datatypes.ClientRequest{Request:tcphttpclient.Start}
	 //clientRequestCh<-myReq

	 select{}
}
