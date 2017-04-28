package main

import (
	"flag"
	"github.com/nearmap/goreportcardlite/handlers"
	"log"
	"net/http"
)

var addr = flag.String("http", ":8000", "HTTP listen address")

func main() {
	flag.Parse()
	http.HandleFunc("/report", handlers.ReportHandler)
	log.Printf("Running on %s ...", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
