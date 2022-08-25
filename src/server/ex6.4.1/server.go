package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

func handler(w http.ResponseWriter, r *http.Request) {
	// dump, err := httputil.DumpRequest(r, true)
	// if err != nil {
	// 	http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(string(dump))
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "#%d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/chunked", handlerChunkedResponse)
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
