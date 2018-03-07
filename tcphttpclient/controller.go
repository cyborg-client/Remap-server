package tcphttpclient

import (
	"fmt"
	"github.com/cyborg-client/client/datatypes"
)

func TcpHttpClient(
	tcpDataStreamCh chan TcpDataStream,
	tcpHttpClientStatusCh chan<- Status,
	clientRequestCh <-chan datatypes.ClientRequest,
) {
	// Channels
	// HTTP -> TCP
	startStopTcpCh := make(chan startStopTcp)

	// Create goroutines
	go httpMain(clientRequestCh, startStopTcpCh, tcpHttpClientStatusCh)
	go tcpMain(tcpDataStreamCh, startStopTcpCh)

	for {
		select {
		case a := <-tcpDataStreamCh:
			fmt.Println(a)

		}
	}
	select {}

}
