package tcphttpclient

import (
	//"fmt"
	"github.com/cyborg-client/client/datatypes"
)

func TcpHttpClient(
	tcpDataStreamCh chan TcpDataStream,
	tcpHttpClientStatusCh chan<- Status,
	clientRequestCh <-chan datatypes.ClientRequest,
) {
	// Channels
	startStopTcpCh := make(chan startStopTcp)

	// Create goroutines
	go httpMain(clientRequestCh, startStopTcpCh, tcpHttpClientStatusCh)
	go tcpMain(tcpDataStreamCh, startStopTcpCh)

	select {}

}
