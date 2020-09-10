package main

import (
	"flag"
	"log"
	"net/http"
)

func loggingHandler(h http.Handler) http.Handler {
	// Basic logging to the console
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {

	port := flag.String("p", "8080", "Port to listen on")
	dir := flag.String("d", ".", "Directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Fatal(http.ListenAndServe(":"+*port, loggingHandler(http.FileServer(http.Dir(*dir)))))
}
