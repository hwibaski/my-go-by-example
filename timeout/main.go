package main

import (
	"fmt"
	"time"
)

// time.After() 함수가 채널을 반환하므로 select 구문을 사용하여 타임아웃을 구현할 수 있습니다널

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// 바로 위의 고루틴은 2초 뒤에 c1 채널에 값을 전달합니다.
	// time.After() 보다 위의 고루틴이 늦게 채널에 값을 전달하므로 select 구문에 걸리지 않고 타임아웃 처리
	select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second): // 1초 뒤에 채널을 반환
			fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
		case res := <-c2:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
	}
}