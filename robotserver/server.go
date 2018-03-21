package robotserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/cyborg-client/client/analysis"
	"github.com/cyborg-client/client/config"
)

var timestampCh <-chan []int64

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
func serverMain(timestampdataChLocal <-chan []int64) {
	// global chan
	timestampCh = timestampdataChLocal
	router := mux.NewRouter()
	router.HandleFunc("/data/{every-ms}", GetData).Methods("GET")
	fmt.Printf("Starting localhost at port: %v. REST: '/data/<millisec>'\n", config.RobotServerPort)
	log.Fatal(http.ListenAndServe(":"+config.RobotServerPort, router))

}
