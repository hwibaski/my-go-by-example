package main

import (
	"fmt"
	"time"
)

// Go의 select 기능을 사용하면 여러 채널 작업을 대기할 수 있습니다.
// 고루틴과 채널을 선택과 결합하는 것은 Go의 강력한 기능입니다.

func main() {
	// 2개의 채널을 생성합니다.
	c1 := make(chan string)
	c2 := make(chan string)

	// c1 채널에 값을 전달하는 고루틴
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	// c2 채널에 값을 전달하는 고루틴
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("loop", i)
		// select 구문은 각 채널에서 값을 기다리고 있습니다.
		// 1초와 3초 Sleep이 동시에 실행되므로 총 실행 시간은 약 3초에 불과합니다.
		select {
			case msg1 := <-c1:
				println("received", msg1)
			case msg2 := <-c2:
				println("received", msg2)
		}
	}
}