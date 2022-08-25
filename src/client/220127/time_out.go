package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	// まず、Contextを作成
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	// Contextを持ったRequestを作成
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:18888", nil)
	// あとは通常通りに実行
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}
