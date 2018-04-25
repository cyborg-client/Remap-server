package tcphttpclient

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cyborg-client/Remap-server/config"
	"github.com/cyborg-client/Remap-server/errorhandling"
	"io"
	"log"
	"net"
	"github.com/cyborg-client/Remap-server/datatypes"
)

// connectTCP Creates an connection to the MEA database, and receives int32s. Puts those into the tcpDataStream channel.
// If a message is received on the stop channel, the function returns after finishing the current batch.
func connectTCP(tcpDataStream chan<- Segment, stop <-chan bool) {
	conn, err := net.Dial("tcp", config.MEAServerAddress+":"+config.MEAServerTcpPort)
	defer conn.Close()
	if err != nil {
		log.Println("Error in tcp.go: ", err)
		errorhandling.Restart()
	}
	fmt.Println("Connected")

	// Allocate a buffer to receive a batch from TCP
	buffer := make([]byte, 60*4*config.SegmentLength)
	for {
		select {
		case <-stop:
			// If we have received a message from the system to stop the connection, we do so here
			// Defer conn.Close() fixes the rest.
			return
		default:
			//
			// Read the data from the TCP stream. 60 electrodes, with SegmentLength number of int32's.
			// Each int32 has four bytes. Thus read 60 * 4 * SegmentLength number of bytes
			//
			lr := io.LimitReader(conn, 60*4*config.SegmentLength)
			_, err := lr.Read(buffer)
			if err != nil {
				log.Println("TCP Connection is too slow, restarting")
				errorhandling.Restart()
			}

			// Data received is in BigEndian, convert this to a int32 array
			var t int32
			var meaSegment [60 * config.SegmentLength]int32

			for i := 0; i < 60*config.SegmentLength; i++ {
				// Read the next four bytes and convert them to binary. Insert the finished binary
				// into the array.
				buf := bytes.NewReader(buffer[(i * 4):(i*4 + 4)])
				err = binary.Read(buf, binary.BigEndian, &t)
				if err != nil {
					log.Println("Error in tcp.go, bianry ready fail: ", err)
					errorhandling.Restart()
				}
				meaSegment[i] = t
			}

			// Pass the finished parsed int32 to the tcpDataStream.
			tcpDataStream <- meaSegment
		}
	}
}
// tcpMain defines the entrypoint for the TCP connection between the client and the MEA server. Requires segmentCh, which it sends
// out received TCP data and startStopTcpCh, which controls when to start and stop the TCP connection.
func tcpMain(
	segmentCh chan<- Segment,
	startStopTcpCh <-chan startStopTcp,
) {
	{
		stopCh := make(chan bool)
		running := false
		for {
			select {
			case s := <-startStopTcpCh:
				// Received a start/stop message. If start => start connectTCP. If stop => send stop message on stopCh
				if s == datatypes.Start && !running {
					go connectTCP(segmentCh, stopCh)
					running = true
				} else if running && s == datatypes.Stop {
					stopCh <- true
					running = false
				}
			}
		}

	}
}
