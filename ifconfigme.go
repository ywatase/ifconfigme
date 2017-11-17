package main
import (
	"fmt"
	"net/http"
	"strings"
	"flag"
    "github.com/fukata/golang-stats-api-handler"
)

func main () {
	var port *int = flag.Int("port", 4978, "listen port")
	var address *string = flag.String("address", "localhost", "listen address")
	flag.Parse()
	http.HandleFunc("/", viewHandler)
    http.HandleFunc("/api/stats", stats_api.Handler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *address, *port), nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var remoteaddr string
    if x_forwarded_for := r.Header["X-Forwarded-For"] ; x_forwarded_for != nil {
		remoteaddr = x_forwarded_for[0]
	} else {
		remoteaddr = r.RemoteAddr
		tmp := strings.Split(remoteaddr, ":")
		remoteaddr = tmp[0]
	}
	fmt.Fprintf(w, "%s\n", remoteaddr)
}

