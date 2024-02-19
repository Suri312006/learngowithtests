package sync

import "sync"

type Counter struct {
    //NOTE: you should not just do sync.Mutex, as it would make the api
    // to unlock outward facing! (think about it dumbass)
    mu sync.Mutex
    // a mutex must not be copied after first use
    count int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    // you want to defer so incase something inside goes wrong, it still gets unlocked
    defer c.mu.Unlock()
    c.count +=1
}

func (c *Counter) Value() int {
	return c.count
}
