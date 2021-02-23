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
	challengeString := r.Header.Get("Izayakun")
	f := fnv.New64()
	f.Write([]byte(challengeString))
	response := f.Sum64()
	w.Header().Set("Shizuchan", fmt.Sprintf("%d", response))
	w.WriteHeader(200)
}
