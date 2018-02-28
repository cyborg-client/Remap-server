package tcphttpclient

import
(
	"github.com/vegardbotnen/EiT-Client/datatypes"
)

func tcpMain(
	statusTcpCh chan<- statusTcp,
	tcpDataStream chan<- datatypes.TcpDataStream,
	startStopTcpCh <-chan startStopTcp,
	){
}
