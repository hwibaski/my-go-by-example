package main

import (
	"fmt"
	"time"
)

// rate limiting은 리소스 사용률을 제어하고 서비스 품질을 유지하기 위한 중요한 메커니즘입니다.
// Go는 고루틴, 채널, 티커를 통해 속도 제한을 우아하게 지원합니다.

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 200ms 간격으로 값을 내보내는 채널 생성
	limiter := time.Tick(time.Millisecond * 200)

	// requests 채널을 close() 했으므로 iteration 가능
	for req := range requests {
		// limiter 채널에서 data를 받을 때까지 블록됨
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("--------------------")
	// ----------------------------------------
	// 전체 속도 제한을 유지하면서 속도 제한 체계에서 짧은 요청을 허용하고 싶을 수 있습니다.
	// 이 경우 리미터 채널을 버퍼링하여 이를 수행할 수 있습니다.
	// 이 버스티리미터 채널은 최대 3개의 이벤트 버스트를 허용합니다.

	// time.Time 타입을 가지는 버퍼가 3개인 채널 생성
	burstyLimiter := make(chan time.Time, 3)

	// 채널에 현재 시간을 3번 보내 채널을 채웁니다.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 200밀리초마다 burstyLimiter 채널에 새로운 시간 값을 보내는 고루틴을 실행합니다.
	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t
		}
	}()

	// 이제 들어오는 요청 5개를 더 시뮬레이션합니다. 이 중 처음 3개는 버스티리미터의 버스트 기능을 활용합니다.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		// burstyLimiter 채널에서 값을 받을 때까지 블록됩니다.
		// 위의 코드에서 burstyLimiter 채널은 3개의 값(time.Now())을 미리 채워놨기 때문에 빠르게 처리됩니다.
		// 그러나 3개의 요청 이후에는 1초 간격으로 요청을 처리합니다후
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}