package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	stats_api "github.com/fukata/golang-stats-api-handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/api/stats", stats_api.Handler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var remoteaddr string
	if x_forwarded_for := r.Header["X-Forwarded-For"]; x_forwarded_for != nil {
		remoteaddr = x_forwarded_for[0]
	} else {
		remoteaddr = r.RemoteAddr
		tmp := strings.Split(remoteaddr, ":")
		remoteaddr = tmp[0]
	}
	fmt.Fprintf(w, "%s\n", remoteaddr)
}
