package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	client := &http.Client{}
	// NOTE: os.Openでとりあえずbodyにつっこんでみる
	file, err := os.Open("photo.jpg")
	request, err := http.NewRequest("POST", "http://localhost:18888", file)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "image/jpeg")
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
