package main

import (
	"fmt"
	"time"
)

// 우리는 종종 미래의 어느 시점에 Go 코드를 실행하거나 일정한 간격을 두고 반복적으로 실행하고 싶을 때가 있습니다.
// Go에 내장된 타이머와 ticker 기능을 사용하면 이 두 가지 작업을 모두 쉽게 수행할 수 있습니다.
// 먼저 타이머를 살펴본 다음 ticker를 살펴보겠습니다.

func main() {
	// 타이머는 미래의 단일 이벤트를 나타냅니다.
	// 타이머에 대기 시간을 지정하면 해당 시간에 알림을 받을 채널이 제공됩니다.
	// 이 타이머는 2초 동안 기다립니다.
	timer1 := time.NewTimer(2 * time.Second) 

	// 타이머가 실행되었음을 나타내는 값을 전송할 때까지 타이머의 채널 C를 <-timer1.C가 block합니다.
	// 즉 2초간 기다린 후에 타이머가 만료되면 타이머가 실행되었음을 나타내는 값을 전송합니다.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 기다리기만 하고 싶었다면 시간.잠자기를 사용할 수도 있습니다.
	// 타이머가 유용한 이유 중 하나는 타이머가 실행되기 전에 타이머를 취소할 수 있기 때문입니다. 다음은 그 예입니다
	// 새로운 타이머 생성
	timer2 := time.NewTimer(time.Second)
	go func() {
		// 타이머가 만료되면 실행됩니다.
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	// 그러나 타이머가 만료되기 전에 타이머를 중지할 수도 있습니다.
	stop2 := timer2.Stop()

	// 타이머가 고루틴에서 만료되기 전에 메인 고루틴에서 중지되었으므로 stop2는 true입니다.
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// timer2가 중지될 수 있는 시간을 주기 위해 메인 고루틴이 2초를 기다린다.
	time.Sleep(2 * time.Second)
}