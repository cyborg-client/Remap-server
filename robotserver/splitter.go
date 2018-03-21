package robotserver

import (
	"github.com/cyborg-client/client/analysis"
)

func splitterMain(timestampdataCh <-chan []int64, registerNewClientCh <-chan clientTimestampCh) {
	activeClients := make([]clientTimestampCh, 100)
	for {
		select {
		case client := <-registerNewClientCh:
			activeClients = append(activeClients, client)
		case msg := <-timestampdataCh:
			ok := false
			for i := 0; i < len(activeClients); i++ {
				activeClients[i], ok <- msg
				if !ok {
					activeClients = append(activeClients[:i], activeClients[i+1:]...)
				}
			}
		}
	}
}

func Main(timestampdataChLocal <-chan []int64) {
	registerNewClientCh := make(chan clientTimestampCh)
	go serverMain(timestampdataChLocal)
	go splitterMain(timestampdataChLocal, registerNewClientCh)
}
