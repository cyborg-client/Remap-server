// Package tcphttpclient implements the TCP and HTTP connection between the client and the MEA server.
package tcphttpclient

import (
	//"fmt"
	"github.com/cyborg-client/client/datatypes"
)

// Main is the entrypoint for the tcphttpclient package. Requires tcpDataStreamCh which it outputs the received TCP data
// and clientRequestCh, which it accepts requests from users.
func Main(
	tcpDataStreamCh chan<- Segment,
	clientRequestCh <-chan datatypes.ClientRequest,
) {
	// Channels
	startStopTcpCh := make(chan startStopTcp)

	// Create goroutines
	go httpMain(clientRequestCh, startStopTcpCh)
	go tcpMain(tcpDataStreamCh, startStopTcpCh)

	select {}

}
