package main

import "fmt"

func main() {
	// 채널은 동시 실행 중인 고루틴을 연결하는 파이프입니다.
	// 한 고루틴에서 채널로 값을 보내고 다른 고루틴에서 해당 값을 받을 수 있습니다.
	messages := make(chan string)

	// 채널에 값을 보내기 위해 <- 연산자를 사용합니다.
	go func() { 
		messages <- "ping"
		fmt.Println("temp")
	}()

	// <- 연산자를 사용하여 채널에서 값을 받습니다.
	// 채널은 수신자와 송신자가 서로를 기다리는 속성 때문에 별다른 동기화 도구 없이 "ping"을 출력할 수 있습니다.
	// 즉 위의 고루틴이 끝날 때 까지 메인 고루틴은 대기하게 됩니다
	msg := <-messages
	fmt.Println(msg)
}