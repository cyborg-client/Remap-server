package robotserver

func splitterMain(timestampdataCh <-chan []int64, registerNewClientCh <-chan splitterRequest, deleteClientCh <-chan splitterRequest) {
	activeClients := make(map[int]splitterRequest)
	for {
		select {
		case client := <-registerNewClientCh:
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
	go serverMain(timestampdataCh)
	go splitterMain(timestampdataCh, registerNewClientCh, deleteClientCh)
	select {}
}
