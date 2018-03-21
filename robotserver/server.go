package robotserver

import (
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
)

var timestampCh <-chan []int64

func ClientHandler(ws *websocket.Conn) {

}

// Main is main func in robotserver-package
func serverMain(timestampdataCh <-chan []int64) {
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(r.URL)
	})
	http.ListenAndServe(":6780", nil)

	select {}
}
