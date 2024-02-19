package selectracer

import (
	//"errors"
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	// myVar := <- ch is a blocking call until the `one` channel is done executing
	// select waits on MULTIPLE channels, and the first onoe to return will be executed
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error){

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <- time.After(timeout)	:
		return "", fmt.Errorf("expected an error but didnt get one")
}
}


// anon struct ??
// creates a channel that is a struct{}
func ping(url string) chan struct{} {
	// in this case, we dont care what type is sent to the channel, as soon as
	// its sent, we can close and return
	// since we are closing instead of sending, we dont need to allocate memory
	ch := make(chan struct{})
	// anon function that is called right after instantiation
	go func() {
		// this is a blocking call, stops execution of the go routine till completed
		http.Get(url)
		// we close the channel, making the funcitono return
		close(ch)
	}()

	// the speed at which we return this function is based off the amount of time
	// it takes the blocking ccall of http.Get() to finish
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
