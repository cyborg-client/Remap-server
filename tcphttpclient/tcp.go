package tcphttpclient

import
(
	"github.com/cyborg-client/client/datatypes"
)

func tcpMain(
	statusTcpCh chan<- statusTcp,
	tcpDataStream chan<- datatypes.TcpDataStream,
	startStopTcpCh <-chan startStopTcp,
	){
}
