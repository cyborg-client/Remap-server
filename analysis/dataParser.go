package analysis

import (
	"fmt"

	"github.com/cyborg-client/client/tcphttpclient"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO: find out how to syncronise all channels
func Main(timeStampChannel chan<- []byte, tcpDataStreamCh <-chan tcphttpclient.TcpDataStream) {
	for {
		/*dat, err := os.Open("/Users/jonasdammen/Development/code/src/github.com/cyborg-client/client/sampledata/2017-10-20_MEA2_100rows_1sec.csv")
		check(err)
		reader := csv.NewReader(bufio.NewReader(dat))

		*/

		var timestampArray= make([][]byte, 0, 10)
		i := 0
		cols := 0
		//reader.Read() // first line is info only
		for {
			//record, err := reader.Read()
			record := <-tcpDataStreamCh
			fmt.Println(record)
			if i == 0 {
				cols = len(record)
			}
			timestampArray = append(timestampArray, make([]byte, cols-1))
			// Stop at EOF.
			/*if err == io.EOF {
				fmt.Println("Breaking due to EOF")
				break
			}*/
			//fmt.Println("reading values")

			for j := range record {
				if j != 0 { // first value is timestamp
					val := record[j]
					if val < 0 {
						timestampArray[i][j-1] = 1

					} else {
						timestampArray[i][j-1] = 0
					}
				}
			}
			/*
		fmt.Println("Record current line")
		fmt.Println((record))
		fmt.Println("TSA current line")
		fmt.Println(timestampArray[i])
		fmt.Println("TSA current length")
		fmt.Println(len(timestampArray))
		*/
			//fmt.Println("Pushing one array to channel")
			//fmt.Println(timestampArray[i])
			timeStampChannel <- timestampArray[i]
			i++
		}
		fmt.Println("Sent ", i, " lines");
		//check(err)	r := csv.NewReader(strings.NewReader(in))
	}
}