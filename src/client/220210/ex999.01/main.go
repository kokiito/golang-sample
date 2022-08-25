package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub() // ここで無限ループに入るので、以下の for 文は実行されない。 go sub() にすると並行処理になる。
	go sub()
	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}

}
