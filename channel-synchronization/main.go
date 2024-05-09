package main

import (
	"fmt"
	"time"
)

// 채널을 사용하여 여러 고루틴의 실행을 동기화할 수 있습니다.
// 다음은 코드 흐름을 block하는 채널 수신을 사용하여 고루틴이 완료될 때까지 기다리는 예시입니다.
// 여러 고루틴이 완료될 때까지 기다릴 때는 WaitGroup을 사용하는 것이 좋습니다.

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 작업이 완료되면 channel에 true를 보냅니다.
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// <-done은 channel로부터 값을 수신합니다.
	// 여기서는 worker 함수가 작업을 완료할 때까지 기다립니다.
	// 이 프로그램에서 <- done 줄을 제거하면 워커가 시작하기도 전에 프로그램이 종료됩니다.
	<-done
}