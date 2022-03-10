package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body><h1>Good evening!</h1></body></html>\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening :18443")
	log.Println(http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil))
}
