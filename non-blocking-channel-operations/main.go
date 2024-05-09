package main

import (
	"fmt"
)

// 채널의 기본 송수신은 blocking 입니다.
// 그러나 default 절과 함께 선택을 사용하여 비차단 송신, 수신 및 비차단 다방향 선택을 구현할 수 있습니다.

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "hi"
	}()

	// message 채널에 값을 보내지 않았으므로 default 절이 실행됩니다.
	// non-blocking receive
	select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")

	}

	// channel에 buffer가 없고, 수신자가 없으므로 default 절이 실행됩니다.
	// non-blocking send
	msg := "hi"
	select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
	}

	// default 위에 여러 case를 사용하여 다방향 논블로킹 select를 구현할 수 있습니다. 
	// 여기서는 message 와 signal 모두에 대해 non-blocking 수신을 시도합니다.
	select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
	}
}