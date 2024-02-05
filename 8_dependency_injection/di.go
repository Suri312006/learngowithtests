package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// `Buffer` type from the bytes package implements the `Writer` interface
// because it has the method Write(p []byte) (n int, err error)
func Greet(writer io.Writer, name string) {
	// we are using buff in our test to send in as our writer to check what
	// is being written to it after function call

	// we essentially broke into how the printf function worked, and injected our own thing to test if its working properly
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
    Greet(w, "world")
}

func main() {

	// Greet(os.Stdout, "Elodie")
    log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
