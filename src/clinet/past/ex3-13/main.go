package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}
	values := url.Values{"test": {"こんにちは！"}}
	reader := strings.NewReader(values.Encode())
	request, err := http.NewRequest("DELETE", "http://localhost:18888", reader)
	request.Header.Add("Content-Types", "image/jpeg")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
