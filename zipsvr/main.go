package main

import (
	"net/http"
	"log"
	"fmt"
	"runtime"
	"encoding/json"
	"github.com/lb7/info344-in-class/zipsvr/models"
	"strings"
	"os"
	"github.com/lb7/info344-in-class/zipsvr/handlers"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello, %s", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(stats)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error loading file: %v", err)
	}
	log.Printf("Loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index: cityIndex,
		PathPrefix: "/zips/",
	}

	mux.Handle("/zips/", cityHandler)

	fmt.Printf("Server is listening at http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
