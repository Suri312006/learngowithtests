package concurrency


import (
)
type WebsiteChecker func(string) bool

type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    // pass channel the struct you want it to use?
    resultChannel := make(chan result)
    // the _ is i, the index
    // each iteration of this loop will start a new goroutine
    for _, url := range urls {
        // this is an anomyous function

        go func(u string) {
            // sending the output to a result channel
            // the <- is a send statement ( channel on left and value on right)
            resultChannel <- result{u, wc(u)}

        }(url)
        // the parenthesis at the end is to be able
        // to call it right after declaration
        // THIS BEHAVIOR IS SPECIFIC TO ANON FUNCTIONS
        
    }

    for i :=0; i<len(urls); i++ {
        // recieve expression (value on left, channel on right)
        r:= <-resultChannel
        results[r.string] = r.bool
    }

    // returns a map of each URL checked to a boolean value:
    // true for a good response, false for a bad response
    return results
}
