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

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		// Check if target directory exists
		log.Fatal("Directory does not exist")
	}

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Fatal(http.ListenAndServe(":"+*port, apachelog.CombinedLog.Wrap(http.FileServer(http.Dir(*dir)), os.Stderr)))

}
