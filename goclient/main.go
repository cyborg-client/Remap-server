package main

import
(
	"./dataTypes"
	"./tcpHttpClient"
)

func main() {
	// Make channels
	/*
	tcpDataStreamCh chan<- dataTypes.TcpHttpClientData,
	tcpHttpClientStatusCh chan<- dataTypes.TcpHttpClientStatus,
	clientRequestCh <-chan dataTypes.ClientHttpServerRequest,
	 */
	 tcpDataStreamCh := make(chan dataTypes.TcpHttpClientData)
	 tcpHttpClientStatusCh := make(chan dataTypes.TcpHttpClientStatus)
	 clientRequestCh := make(chan dataTypes.ClientHttpServerRequest)

	 go tcpHttpClient.TcpHttpClientMain(tcpDataStreamCh, tcpHttpClientStatusCh, clientRequestCh)

	 select{}
}
