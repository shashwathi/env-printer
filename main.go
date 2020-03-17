package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("Starting application")

	port := flag.Int("port", 8080, "port to listen on")

	flag.Parse()

	http.HandleFunc("/", printRequest)

	listenAddr := fmt.Sprintf("%s:%d", "", *port)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		panic(err)
	}
}

func printRequest(w http.ResponseWriter, r *http.Request) {
	var output []string
	output = append(output, fmt.Sprintf("POD_NAME: %s", os.Getenv("POD_NAME")))
	result := strings.Join(output, "\n") + "\n"
	fmt.Fprintf(w, result)
}
