package main

import (
	"flag"
	apachelog "github.com/lestrrat-go/apache-logformat"
	"log"
	"net/http"
	"os"
)

func main() {

	port := flag.String("p", "8080", "Port to listen on")
	dir := flag.String("d", ".", "Directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Fatal(http.ListenAndServe(":"+*port, apachelog.CombinedLog.Wrap(http.FileServer(http.Dir(*dir)), os.Stderr)))

}
