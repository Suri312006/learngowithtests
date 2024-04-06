package main

import (
    "log"
    "net/http"
 "github.com/suri312006/learnGoWithTests/http_server/server"   
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return 123
}


func main(){
    server := &httpserver.PlayerServer{&InMemoryPlayerStore{}}
    log.Fatal(http.ListenAndServe(":5000", server))
}
