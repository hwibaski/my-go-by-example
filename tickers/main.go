package main

import (
	"fmt"
	"time"
)

// 타이머는 미래에 한 번만 하고 싶은 일을 할 때 사용하는 반면, 티커는 일정한 간격으로 반복적으로 하고 싶은 일을 할 때 사용합니다.
// 다음은 중지할 때까지 주기적으로 틱하는 티커의 예입니다.

func main() {
	// 500 밀리초마다 틱을 받는 새로운 타이머를 생성합니다.
	ticker := time.NewTicker(500 * time.Millisecond)
	// done 채널을 생성합니다.
	done := make(chan bool)

	go func() {
		for {
			select {
				// done 채널에 true가 전달되면 루프를 종료합니다에
				case <-done:
					return
				// ticker.C 채널에서 값을 받으면 틱을 출력합니다.
				case t := <-ticker.C:
					fmt.Println("Tick at", t)
			}
		}
	}()

	// 1600 밀리초 동안 대기한 후 타이머를 중지합니다.
	time.Sleep(1600 * time.Millisecond)
	// ticker를 중지합니다.
	ticker.Stop()
	// 무한 고루틴을 종료시키기 위해 done 채널에 true를 보냅니다.
	done <- true
	fmt.Println("Ticker stopped")
}