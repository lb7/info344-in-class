package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"net/http/httputil"
	"sync"
)

//RootHandler handles requests for the root resource
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello from the gateway! Try requesting /v1/time")
}

func NewServiceProxy(addrs []string) *httputil.ReverseProxy {
	nextIndex := 0
	mx := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			//modify request to indicate remote host
			mx.Lock()
			r.URL.Host = addrs[nextIndex % len(addrs)]
			nextIndex++
			mx.Unlock()

			r.URL.Scheme = "http"
		},
	}
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	//TODO: get network addresses for our
	//timesvc instances
	timesvcAddrs := os.Getenv("TIMESVC_ADDRS")
	splitTimeSvcAddrs := strings.Split(timesvcAddrs, ",")

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	//TODO: add reverse proxy handler for `/v1/time`
	mux.Handle("/v1/time", NewServiceProxy(splitTimeSvcAddrs))

	log.Printf("server is listening at https://%s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "tls/fullchain.pem", "tls/privkey.pem", mux))
}