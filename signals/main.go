package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 때때로 우리는 Go 프로그램이 유닉스 신호를 지능적으로 처리하기를 원합니다.
// 예를 들어, 서버가 SIGTERM을 수신하면 서버가 정상적으로 종료되거나 명령줄 도구가 SIGINT를 수신하면 입력 처리를 중지하기를 원할 수 있습니다.
// 다음은 채널로 Go에서 신호를 처리하는 방법입니다.
func main() {
	// 이동 신호 알림은 채널에서 os.Signal 값을 전송하는 방식으로 작동합니다.
	// 이러한 알림을 수신할 채널을 만들겠습니다.
	// 이 채널은 버퍼링되어야 한다는 점에 유의하세요.
	sigs := make(chan os.Signal, 1)

	// signal.Notify는 지정된 채널을 등록하여 지정된 신호에 대한 알림을 수신합니다.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 여기서는 메인 함수에서 시그로부터 수신할 수 있지만, 보다 현실적인 정상 종료 시나리오를 보여주기 위해 별도의 고루틴에서도 이 작업을 수행할 수 있는 방법을 살펴보겠습니다.
	done := make(chan bool, 1)

	go func() {
		// 이 루프는 우리가 SIGINT나 SIGTERM을 수신할 때까지 기다립니다.
		sig := <-sigs
		fmt.Println()
		// 수신한 신호를 출력하고 done 채널에 true를 보냅니다.
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	// done 채널로부터 값을 받을 때까지 블로킹됩니다.
	<-done
	fmt.Println("exiting")
}
