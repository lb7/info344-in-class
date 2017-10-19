package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultAddr = ":80"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	//TODO: make connection to the DBMS
	//construct the appropriate tasks.Store
	//construct the handlers.Context

	mux := http.NewServeMux()
	//mux.HandleFunc("/v1/tasks", TODO: add TasksHandler )
	//mux.HandleFunc("/v1/tasks/", TODO: add SpecificTaskHandler )

	fmt.Printf("server is listening at http://%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
