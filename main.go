package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting application on PORT 9001")
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

// HelloServer responds with a simple static message
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Server")
}
