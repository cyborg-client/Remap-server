package websocketserver

import (
	"github.com/cyborg-client/client/errorhandling"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"github.com/cyborg-client/client/config"
)

// splitterMain is a PUBSUB implementation, taking the data from timestampdataCh and broadcasting it to all subscribers. In order
// to register as a subscriber, you must register yourself by sending a request on the registerNewClientCh. When you do no longer wish to
// subscribe, you must delete yourself by sending a request on the deleteClientCh.
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
// Main is the entry point for the websocketserver. Requires timestampdataCh, which expects parsed data from the analysis package.
func Main(parsedDataCh <-chan []int64) {
	registerNewClientCh := make(chan splitterRequest)
	deleteClientCh := make(chan splitterRequest)
	go serverMain(registerNewClientCh, deleteClientCh)
	go splitterMain(parsedDataCh, registerNewClientCh, deleteClientCh)
	http.HandleFunc("/stimulate", stimulateServer)
	log.Fatal(http.ListenAndServe(":"+ config.StimulateServerPort, nil))
	select {}
}
