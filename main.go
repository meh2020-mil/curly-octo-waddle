package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":80", h{})
}

type h struct {
}

func (h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Headers", r.Header)
	challengeString := r.Header.Get("CHALLENGE")
	f := fnv.New64()
	f.Write([]byte(challengeString))
	response := f.Sum64()
	w.Header().Set("RESPONSE", fmt.Sprintf("%d", response))
	w.WriteHeader(200)
}
