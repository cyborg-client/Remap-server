package tcphttpclient

import
(
	"../dataTypes"
)

func tcpMain(
	statusTcpCh chan<- statusTcp,
	tcpDataStream chan<- dataTypes.TcpDataStream,
	startStopTcpCh <-chan startStopTcp,
	){
}
