package buffer

import "fmt"

// Main is the main function in buffer package
func Main(timeStampChannel <-chan []byte) {
	fmt.Println("dmnlk")
	counter := 0;
	for array := range timeStampChannel {
		fmt.Println("buffer ", array)
		counter += 1;
		fmt.Println("Recieved ", counter, " lines");
	}
}
