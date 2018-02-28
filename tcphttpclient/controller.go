package tcphttpclient

import(
	"../dataTypes"
)

func TcpHttpClientMain(
	tcpDataStreamCh chan<- dataTypes.TcpDataStream,
	tcpHttpClientStatusCh chan<- dataTypes.TcpHttpClientStatus,
	clientRequestCh <-chan dataTypes.ClientRequest,
	) {
		// Channels
		// HTTP -> TCP
		startStopTcpCh := make(chan startStopTcp)
		// TCP -> HTTP
		statusTcpCh := make(chan statusTcp)

		// Create goroutines
		go httpMain(statusTcpCh, clientRequestCh, startStopTcpCh, tcpHttpClientStatusCh)
		//go tcpMain(statusTcpCh, startStopTcpCh)
}
