package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto" // 追加
	"os"
)

func main() {
	var buffer bytes.Buffer                      // ❶
	writer := multipart.NewWriter(&buffer)       // ❷
	writer.WriteField("name", "Michael Jackson") // ❸
	// ❹
	// 修正箇所ここから
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg") // ❺
	if err != nil {
		// ファイル読み込み失敗
		panic(err)
	}
	// 修正箇所ここまで
	io.Copy(fileWriter, readFile) // ❻
	writer.Close()                // ❼
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}