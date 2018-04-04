package robotserver

import (
	"github.com/cyborg-client/client/errorhandling"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"github.com/cyborg-client/client/config"
)

func splitterMain(timestampdataCh <-chan []int64, registerNewClientCh <-chan splitterRequest, deleteClientCh <-chan splitterRequest) {
	activeClients := make(map[uuid.UUID]splitterRequest)
	for {
		select {
		case client := <-registerNewClientCh:
			if _, ok := activeClients[client.ID]; ok {
				log.Println("Error: UUID already exists")
				errorhandling.Restart()
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
	http.HandleFunc("/stimulate", stimulateServer)
	log.Fatal(http.ListenAndServe(":"+ config.StimulateServerPort, nil))
	select {}
}
