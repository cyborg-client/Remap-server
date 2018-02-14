package tcpHttpClient

import(
	"../dataTypes"
	"fmt"
)


func httpMain(
	statusTcpCh <-chan  httpTcpStatusMessage,
	startStopTcpCh chan<- tcpHttpStatus,
	clientHttpServerRequest <-chan dataTypes.ClientHttpServerRequest,
	){
		select{
			case req := <- clientHttpServerRequest:
				if req.Request == dataTypes.Start {
					fmt.Println("start")
				}else if req.Request == dataTypes.Stop {
					fmt.Println("Stop")
				}
		}
}