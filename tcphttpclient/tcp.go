package tcphttpclient

import (
	"fmt"
	"net"
	"github.com/cyborg-client/client/config"
	"os"
	"io"
)

func connectTCP(tcpDataStream chan<- TcpDataStream, stop <-chan bool) {
	fmt.Println("Running")
	conn, err := net.Dial("tcp", config.MEAServerAddress+":"+config.MEAServerTcpPort)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(22)
	}
	fmt.Println("Connected")
	b := make([]byte, 60*config.SampleRate*4)
	for {
		select {
		case <-stop:
			return
		default:
			lr := io.LimitReader(conn, 60*config.SampleRate*4)
			status, err := lr.Read(b)
			if err != nil {
				fmt.Println(err)
				os.Exit(42)
			}
			fmt.Println(b)
			if err != nil {
				// HANDLE
				fmt.Println(err)
				os.Exit(44)
			}
			fmt.Println(status)
		}
	}
}

func tcpMain(
	statusTcpCh chan<- statusTcp,
	tcpDataStream chan<- TcpDataStream,
	startStopTcpCh <-chan startStopTcp,
) {
	{
		stopCh := make(chan bool)
		running := false
		select {
		case s := <-startStopTcpCh:
			if s == Start {
				if (!running){
					go connectTCP(tcpDataStream, stopCh)
					running = true
				}
			}else {
				stopCh <- true
				running = false
			}
		}
	}
}
