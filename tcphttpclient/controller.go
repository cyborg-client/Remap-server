package tcphttpclient

import(
	"github.com/vegardbotnen/EiT-Client/datatypes"
)

func TcpHttpClientMain(
	tcpDataStreamCh chan<- datatypes.TcpDataStream,
	tcpHttpClientStatusCh chan<- datatypes.TcpHttpClientStatus,
	clientRequestCh <-chan datatypes.ClientRequest,
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
