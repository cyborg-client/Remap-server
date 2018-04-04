package robotserver

import (
	"encoding/json"
	"fmt"
	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/config"
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
	"net/http"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"
)

// sendAndResetBucket sends the bucket array out on the websocket connection.
// After sending the bucket onto the websocket, it sets all the element in the buffer
// to zero, and sets the timeNow the last recorded time in analysis.
func sendAndResetBucket(bucket *[]int64, ws *websocket.Conn, timeNow *int64) {
	b, _ := json.Marshal(bucket)
	var jsonStr = []byte(b)
	ws.Write(jsonStr)

	for i := range *bucket {
		(*bucket)[i] = 0

	}
	*timeNow = atomic.LoadInt64(&analysis.TimeStamp)
}

// clientHandler handles a websocket connection. It processes the timestampdata from the analysis module,
// and puts them into their respective 60 buckets. When everyMs has passed, it calls the sendAndResetBucket function
func clientHandler(ws *websocket.Conn, everyMs int64, dataCh <-chan analysis.Timestampdata) {
	connectionClosedCh := make(chan bool)
	go func(connectionClosedCh chan bool) {
		_, err := ws.Read(make([]byte, 1))
		if err != nil {
			fmt.Println(err)
			connectionClosedCh <- true
		}
	}(connectionClosedCh)

	bucket := make([]int64, 60)
	var timeNow int64
	timeNow = atomic.LoadInt64(&analysis.TimeStamp)
	for {
		// Connection closed
		select {
		case msg := <-dataCh:
			bucket[msg[1]] += 1
			if (msg[0] - timeNow) > (everyMs * 1000) {
				sendAndResetBucket(&bucket, ws, &timeNow)
			}
		case <-connectionClosedCh:
			return
		case <-time.After(time.Millisecond):
			if (atomic.LoadInt64(&analysis.TimeStamp) - timeNow) > (everyMs * 1000) {
				sendAndResetBucket(&bucket, ws, &timeNow)
			}
		}
	}
}

// serverMain is the entrypoint for the websocket server. Takes as input a channel with timestamp data from the
// analysis module, two splitterRequest channels registering and deleting itself
func serverMain(timestampdataCh <-chan []int64, registerNewClientCh chan<- splitterRequest, deleteClientCh chan<- splitterRequest) {
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(r.URL)
		// Get ID as variable
		re := regexp.MustCompile("^/data/(\\d+)/?$")
		regexParsed := re.FindStringSubmatch(r.URL.String())
		if len(regexParsed) != 2 {
			fmt.Println("Returning error")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ms, err := strconv.Atoi(regexParsed[1])
		if err != nil {
			fmt.Println("Returning error")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Register as listener
		u1 := uuid.Must(uuid.NewV4())
		dataCh := make(chan analysis.Timestampdata, 100)
		clientStruct := splitterRequest{u1, dataCh}
		registerNewClientCh <- clientStruct
		websocket.Handler.ServeHTTP(func(ws *websocket.Conn) {
			clientHandler(ws, int64(ms), dataCh)
		}, w, r)
		deleteClientCh <- clientStruct
	})
	http.ListenAndServe(":"+config.WebSocketPort, nil)

	select {}
}
