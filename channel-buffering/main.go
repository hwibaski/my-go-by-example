package main

import "fmt"

func main() {
	// 기본적으로 채널은 버퍼링되지 않은 상태로, 전송된 값을 받을 준비가 된 수신(<- 채널)이 있는 경우에만 전송(chan <-)을 수락한다는 의미입니다.
	// 버퍼링된 채널은 해당 값에 해당하는 수신자가 없는 경우 제한된 수의 값만 허용합니다.
	messages := make(chan string, 2)

	// 이 채널은 버퍼링되어 있으므로 해수신 없이도 이러한 값을 채널로 보낼 수 있습니다.
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func sample1() {
	// 채널에 버퍼가 없는 상태
	messages := make(chan string)

	// deadlock 발생
	messages <- "buffered"
	messages <- "channel"
}

func sample2() {
	// 채널에 버퍼가 없는 상태
	messages := make(chan string)

	messages <- "buffered"

	// 송신자만 있고 수신자가 없는 상태에서는 deadlock 발생
}