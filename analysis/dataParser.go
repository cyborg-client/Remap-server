package analysis

import (
	"github.com/cyborg-client/client/tcphttpclient"
	"time"
)

func Main(timeStampChannel chan<- []int64, tcpDataStreamCh <-chan tcphttpclient.TcpDataStream) {
	for {
		var timestampTuple = make([]int64, 0, 2)
		var start time.Time
		var effect float64
		var average float64
		var threshold float64
		var MEAChannel int64
		var timeStamp int64
		average = 0
		effect = 0.1
		threshold = 5000000
		start = time.Now()

		for {
			record := <-tcpDataStreamCh
			for j := range record {
				val := record[j]
				// UPDATE FILTER
				average = (1 - effect) * average + effect * float64(-val)
				diff := float64(val) - average

				// SEND TIMESTAMP
				if diff > threshold {
					timestampTuple = make([]int64, 0, 2)
					MEAChannel = int64(j)
					timeStamp = int64(time.Now().Sub(start)) //int64
					timestampTuple = append(timestampTuple, timeStamp)
					timestampTuple = append(timestampTuple, MEAChannel)
					timeStampChannel <- timestampTuple
				}
			}
		}
	}
}