package main

// 채널을 함수 매개변수로 사용할 때 채널이 값만 전송할 것인지 아니면 수신할 것인지 지정할 수 있습니다.
// 이러한 특정성은 프로그램의 유형 안전성을 높입니다.

// chan<- : 이 채널은 값을 전송만 할 수 있습니다.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// <-chan : 이 채널은 값을 수신만 할 수 있습니다.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	println(<-pongs)
}