package main

import (
	"fmt"
	"time"
)

func BName() {
	arr1 := [4]string{"aa", "bb", "cc", "dd"}
	for t1 := 0; t1 < 1000; t1++ {
		// time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1%4])
	}
}

func BId() {
	arr2 := [...]int{
		11, 22, 33, 44,
	}

	for t2 := 0; t2 < 100; t2++ {
		// time.Sleep(150 * time.Microsecond)
		fmt.Printf("%d\n", arr2[t2%4])
	}
}

func f(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s, ":", i)
	}
}

func channelTest(ch chan int) {
	fmt.Println("开始测试协程阻塞")
	time.Sleep(time.Second)
	fmt.Println("协程阻塞结束")
	ch <- 0
	time.Sleep(time.Second)
	fmt.Println("协程结束")
}

// func main() {
// 	ch := make(chan int)
// 	fmt.Println(runtime.NumCPU())
// 	fmt.Println("! ----- 主协程开始 ------")

// 	go BName()
// 	go BId()

// 	go f("direct")
// 	go f("goroutine")
// 	go channelTest(ch)
// 	<-ch
// 	// time.Sleep(3500 * time.Millisecond)

// 	fmt.Println("!------ 主协程结束 ------")
// }
