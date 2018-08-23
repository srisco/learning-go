package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom type implementing the Writer interface
type logWritter struct{}

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Make a new byte slice with space for 99999 elements
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// Using the io.Copy function to print out the body
	// io.Copy(os.Stdout, res.Body)

	// Using the ioutil.ReadAll function (don't need to create the byte slice manually)
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(body))

	// Using io.Copy with our custom Write type
	lw := logWritter{}
	io.Copy(lw, res.Body)
}

// Method Write to implement the interface
// Important to implement according to the specification
// (Go doesn't care about that)
func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("\nJust wrote this many bytes:", len(bs))

	return len(bs), nil
}
