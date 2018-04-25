// Package analysis implements the parsing of raw MEA data into spike data.
package analysis

import (
	"github.com/cyborg-client/Remap-server/tcphttpclient"
	"sync/atomic"
)

// TimeStamp is a global variable used to synchronize the latest between modules. Must be read and written with
// atomic.LoadInt64() and atomic.AddInt64()
var TimeStamp int64

// Main implements the analysis loop. Receives data from tcphttpclient, create the spike data and sends them to the
// websocket server. In practice, it implements a high-pass filter, using a floating cutoff.
func Main(timeStampChannel chan<- []int64, tcpDataStreamCh <-chan tcphttpclient.Segment) {
	for {
		var timestampTuple = make([]int64, 0, 2)
		var effect float64
		var average float64
		var threshold float64
		var MEAChannel int64
		//	average = 0
		effect = 0.1
		threshold = 5000000
		TimeStamp = 0

		var wasActive = make([]int8, 60, 60)
		var averages = make([]float64, 60, 60)

		for {
			record := <-tcpDataStreamCh
			atomic.AddInt64(&TimeStamp, 100)
			for j := range record {
				val := -record[j]
				// UPDATE FILTER
				averages[j] = (1 - effect) * averages[j] + effect * float64(val) // UPDATES FILTER
				diff := float64(val) - averages[j]

				// SEND TIMESTAMP
				if diff > threshold && wasActive[j] == 0 {
					wasActive[j] = 1
					timestampTuple = make([]int64, 0, 2)
					MEAChannel = int64(j)
					timestampTuple = append(timestampTuple, atomic.LoadInt64(&TimeStamp))
					timestampTuple = append(timestampTuple, MEAChannel)
					timeStampChannel <- timestampTuple
				}
				if wasActive[j] == 1 {
					wasActive[j] = 0
				}
			}
		}
	}
}
