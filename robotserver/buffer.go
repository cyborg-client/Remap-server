package robotserver

import (
	"fmt"
	"net/http"
	//"time"
	"time"
	"github.com/cyborg-client/client/analysis"
)

// Main is the main function in buffer package
func GetChunks(everyMS int, w http.ResponseWriter) {
	fmt.Println("Buffer starts creating chunks.")
	timestampCh := make(chan analysis.Timestampdata) // TODO: RENAME

	// buffer size
	size := everyMS * 10
	fmt.Println("Buffer size:", size)

	// allocating buffer
	buffer := make([]int32, 60)

	// fill buffer
	arr := make([]byte, 60)
	for {
		//time.Sleep(time.Until()) // sleepUntilNextIteration : everyMS
		for t := 0; t < size; t++ {
			// fmt.Printf("Buff [%v]: %v\n", t, buffer)
			arr = <-timestampCh
			// fmt.Println("Length..", len(serverCh))
			for i := 0; i < 60; i++ {
				buffer[i] += int32(arr[i])
			}
		}

		// send buffer to receiver
		fmt.Fprintf(w, "[")
		for i, v := range buffer{
			if i != len(buffer){
				fmt.Fprintf(w, "%d,", v)
			}else{
				fmt.Fprintf(w, "%d]\n", v)
			}
		}

		//time.Sleep(time.Second)
		//os.Exit(23)
		//fmt.Fprintf(w, "%v", []byte(string(buffer)))

		// zeroing
		arr = nil
		arr = arr[:cap(arr)]
	}
}