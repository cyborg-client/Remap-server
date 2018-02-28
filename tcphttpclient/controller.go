package tcphttpclient

import(
	"github.com/cyborg-client/client/datatypes"
)

func TcpHttpClientMain(
	tcpDataStreamCh chan<- TcpDataStream,
	tcpHttpClientStatusCh chan<- TcpHttpClientStatus,
	clientRequestCh <-chan datatypes.ClientRequest,
	) {
		// Channels
		// HTTP -> TCP
		startStopTcpCh := make(chan startStopTcp)
		// TCP -> HTTP
		statusTcpCh := make(chan statusTcp)

		// Create goroutines
		//go httpMain(statusTcpCh, clientRequestCh, startStopTcpCh, tcpHttpClientStatusCh)
		go tcpMain(statusTcpCh, tcpDataStreamCh, startStopTcpCh)
		startStopTcpCh <- Start

}
