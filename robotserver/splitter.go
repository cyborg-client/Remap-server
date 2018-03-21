package robotserver

import (
	"github.com/satori/go.uuid"
)

func splitterMain(timestampdataCh <-chan []int64, registerNewClientCh <-chan splitterRequest, deleteClientCh <-chan splitterRequest) {
	activeClients := make(map[uuid.UUID]splitterRequest)
	for {
		select {
		case client := <-registerNewClientCh:
			if _, ok := activeClients[client.ID]; ok {
				panic("UUID already exists")
			}
			activeClients[client.ID] = client
		case msg := <-timestampdataCh:
			for _, v := range activeClients {
				v.DataCh <- msg
			}
		case client := <-deleteClientCh:
			delete(activeClients, client.ID)
		}
	}
}

func Main(timestampdataCh <-chan []int64) {
	registerNewClientCh := make(chan splitterRequest)
	deleteClientCh := make(chan splitterRequest)
	go serverMain(timestampdataCh, registerNewClientCh, deleteClientCh)
	go splitterMain(timestampdataCh, registerNewClientCh, deleteClientCh)
	select {}
}
