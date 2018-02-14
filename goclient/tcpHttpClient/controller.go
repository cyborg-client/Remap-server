package tcpHttpClient

import(
	"../dataTypes"
)

func TcpHttpClientMain(
	tcpDataStreamCh chan<- dataTypes.TcpHttpClientData,
	tcpHttpClientStatusCh chan<- dataTypes.TcpHttpClientStatus,
	clientRequestCh <-chan dataTypes.ClientHttpServerRequest,
	) {
		// Channels
		// HTTP -> TCP
		startStopTcpCh := make(chan tcpHttpStatus)
		// TCP -> HTTP
		statusTcpCh := make(chan httpTcpStatusMessage)

		// Create goroutines
		go httpMain(statusTcpCh, startStopTcpCh, clientRequestCh)
		go tcpMain(statusTcpCh, startStopTcpCh)
}
