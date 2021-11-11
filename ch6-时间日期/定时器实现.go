package main

import (
	"fmt"
	"time"
)

// 定时器的实现
// time.Sleep, NewTimer, AfterFunc

func TestAfter() {
	// time.After
	c := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		<-c
	}()
	select {
	case c <- 1:
		fmt.Println("channel")
	case <-time.After(2 * time.Second):
		// 2秒后关闭channel
		close(c)
		fmt.Println("timeout")
	}
}

// 停止定时器或者重启定时器
func TestTimer() {
	start := time.Now()

	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after 2 second", time.Now().Sub(start))
	})
	time.Sleep(1 * time.Second)
	// reset在 time未触发时返回true, 出发后返回false
	if timer.Reset(3 * time.Second) {
		fmt.Println("time is not trigger")
	} else {
		fmt.Println("timer had expired or stop!")
	}
	time.Sleep(10 * time.Second)
	fmt.Print(start)
}

func main() {
	TestTimer()
}
