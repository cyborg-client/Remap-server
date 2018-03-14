package robotserver

import (
	"fmt"
	"net/http"
)

// Main is the main function in buffer package
func GetChunks(everyMS int, w http.ResponseWriter) {
	fmt.Println("Buffer starts creating chunks.")

	// buffer size
	size := everyMS * 10
	fmt.Println("Buffer size:", size)

	// allocating buffer
	buffer := make([]int32, 60)

	// fill buffer
	arr := make([]byte, 60)
	for {
		for t := 0; t < size; t++ {
			// fmt.Printf("Buff [%v]: %v\n", t, buffer)
			arr = <-timestampCh
			// fmt.Println("Length..", len(serverCh))
			for i := 0; i < 60; i++ {
				buffer[i] += int32(arr[i])
			}
		}

		// send buffer to receiver
		fmt.Fprintf(w, "%v", []byte(string(buffer)))

		// zeroing
		arr = nil
		arr = arr[:cap(arr)]
	}
}