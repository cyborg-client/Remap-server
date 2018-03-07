package robotserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var serverCh <-chan []byte

// GetData parses REST request
func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range params {
		// if param exist, check whether it's an integer
		everyMS, err := strconv.Atoi(params["every-ms"])
		if err != nil {
			fmt.Printf("Error! Invalid 'every-ms' type, should be integer, found: '%v'\n", item)
		} else {
			// send request to buffer-package
			// TODO: call buffer.GetDataEveryMs(everyMs, ...)
			fmt.Printf("Receiving data request for every %v ms\n", everyMS)
			// request data from buffer
			GetChunks(everyMS, w)
		}
	}
}

// Main is main func in robotserver-package
func Main(port string, ch <-chan []byte) {
	// global chan
	serverCh = ch
	router := mux.NewRouter()
	router.HandleFunc("/data/{every-ms}", GetData).Methods("GET")
	fmt.Printf("Starting localhost at port: %v. REST: '/data/<millisec>'\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
