# Concurrency

#### having more than one thing in progress

### defenitions
- blocking
    - we are waiting for this fuction to finish, before 
executing the next one 

- goroutines

    - non blocking, and will run in a seperate process called a goroutine
    - turn a function into a goroutine by putting `go` infront of the func declaration


### Race conditions
- a bug that occurs when the output of our software is dependent on the timing
and sequence of events that we have no control over

- Go has a built in _race detector_ 
- go test -race


### Channels

- Channels are a Go data structure that can both recieve and send values
- allow communication between different processes
