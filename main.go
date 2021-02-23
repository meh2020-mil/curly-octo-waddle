package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
)

func main() {
	http.ListenAndServe("::8080", h{})
}

type h struct {
}

func (h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	challengeString := r.Header.Get("CHALLENGE")
	f := fnv.New64()
	f.Write([]byte(challengeString))
	response := f.Sum64()
	w.Header().Set("RESPONSE", fmt.Sprintf("%d", response))
	w.WriteHeader(200)
}
