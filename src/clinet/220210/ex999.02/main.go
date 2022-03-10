package main

import "fmt"

func main() {
	ch0 := make(chan int) // channel は make 関数で初期化

	// 受信専用
	// ch0 := make(<-chan int)

	// 送信専用
	// ch0 := make(chan<- int)

	fmt.Printf("cap(ch0) = %v\n", cap(ch0)) // デフォルトでは cap が 0
	fmt.Printf("len(ch0) = %v\n", len(ch0)) // 0

	ch1 := make(chan int, 3)                // cap 3 のキュー構造
	fmt.Printf("cap(ch1) = %v\n", cap(ch1)) // 3
	fmt.Printf("len(ch1) = %v\n", len(ch1)) // 0

	// ch0 <- 1   // enqueue (cap が 0 だとここで deadlock)
	// i := <-ch0 // dequeue (len が 0 だとここで deadlock)
	// fmt.Println(i)

	ch1 <- 100                              // enqueue
	fmt.Printf("cap(ch1) = %v\n", cap(ch1)) // 3
	fmt.Printf("len(ch1) = %v\n", len(ch1)) // 1

	i1 := <-ch1                               // dequeue
	fmt.Printf("dequeue from ch1 = %v\n", i1) // 100
	fmt.Printf("cap(ch1) = %v\n", cap(ch1))   // 3
	fmt.Printf("len(ch1) = %v\n", len(ch1))   // 0

	// i2 := <-ch1 // dequeue (len が 0 だとここで deadlock)
	// fmt.Printf("dequeue from ch2 = %v\n", i2)

	ch1 <- 100                              // enqueue
	ch1 <- 200                              // enqueue
	ch1 <- 300                              // enqueue
	fmt.Printf("cap(ch1) = %v\n", cap(ch1)) // 3
	fmt.Printf("len(ch1) = %v\n", len(ch1)) // 3

	// ch1 <- 400         // enqueue (cap - len が 0 だとここで deadlock)
	fmt.Println(<-ch1) // dequeue
	fmt.Println(<-ch1) // dequeue
	fmt.Println(<-ch1) // dequeue

	fmt.Println(<-ch1) // dequeue (len が 0 だとここで deadlock)
}
