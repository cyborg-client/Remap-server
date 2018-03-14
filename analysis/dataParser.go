package analysis

import (
	"fmt"
	"github.com/cyborg-client/client/tcphttpclient"
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO: find out how to syncronise all channels
func Main(timeStampChannel chan<- []byte) {
	dat, err := os.Open("./sampledata/2017-10-20_MEA2_100rows_1sec.csv")
	check(err)
	reader := csv.NewReader(bufio.NewReader(dat))
	timestampArray := make([][]byte, 0, 100)
	reader.Read() // first line is info only

	// read file
	lc := 0
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		timestampArray = append(timestampArray, make([]byte, 60))
		for i := 0; i < 60; i++ {
			val, _ := strconv.Atoi(record[i])
			if val < 0 {
				timestampArray[lc][i] = byte(1)
			} else {
				timestampArray[lc][i] = byte(0)
			}
		}
		lc++
	}

	for {
		for row := 0; row < 100; row++ {
			timeStampChannel <- timestampArray[row]
		}
	}
}
