package tcphttpclient

import (
	//"fmt"
	"github.com/cyborg-client/client/datatypes"
)

func Main(
	tcpDataStreamCh chan TcpDataStream,
	clientRequestCh <-chan datatypes.ClientRequest,
) {
	// Channels
	startStopTcpCh := make(chan startStopTcp)

	// Create goroutines
	go httpMain(clientRequestCh, startStopTcpCh)
	go tcpMain(tcpDataStreamCh, startStopTcpCh)

	select {}

}
