package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// goroutine은 호출하는 함수와 별도로 실행된다.
	go f("goroutine")

	// 익명함수를 사용하여 goroutine을 실행한다
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 이제 두 함수 호출이 별도의 고루틴에서 비동기적으로 실행되고 있습니다.
	// 두 함수가 완료될 때까지 기다립니다(보다 강력한 접근 방식을 원한다면 WaitGroup을 사용하세요).
	time.Sleep(time.Second)
	fmt.Println("done")
}