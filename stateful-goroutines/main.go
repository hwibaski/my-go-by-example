package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 이전 예제에서는 뮤텍스와 함께 명시적 잠금을 사용하여 여러 고루틴에서 공유 상태에 대한 액세스를 동기화했습니다.
// 또 다른 옵션은 고루틴과 채널에 내장된 동기화 기능을 사용해 동일한 결과를 얻는 것입니다.
// 이 채널 기반 접근 방식은 정확히 하나의 고루틴이 각 데이터를 소유하고 통신함으로써 메모리를 공유한다는 Go의 아이디어와 일치합니다.

// 이 예제에서는 상태가 단일 고루틴에 의해 소유됩니다.
// 이렇게 하면 동시 액세스로 인해 데이터가 손상되지 않습니다.
// 해당 상태를 읽거나 쓰기 위해 다른 고루틴은 소유 고루틴에 메시지를 보내고 그에 상응하는 응답을 받습니다.
// readOp 및 writeOp 구조는 해당 요청을 캡슐화하고 소유 고루틴이 응답하는 방법을 제공합니다.
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 읽기 및 쓰기 작업의 수를 추적하기 위해 두 변수를 선언합니다.
	var readOps uint64
	var writeOps uint64

	// 읽기 및 쓰기 채널은 다른 고루틴에서 각각 읽기 및 쓰기 요청을 발행하는 데 사용됩니다.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 다음은 상태를 소유하는 고루틴으로, 이전 예제에서와 같이 맵이지만 이제는 상태 저장 고루틴에 비공개입니다.
	// 이 고루틴은 읽기 및 쓰기 채널을 반복적으로 선택해 요청이 도착하면 응답합니다. 
	// 응답은 먼저 요청된 작업을 수행한 다음 응답 채널에 성공(읽기의 경우 원하는 값)을 나타내는 값을 전송하는 방식으로 실행됩니다.
	go func() {
		// 상태를 저장하기 위한 맵
		var state = make(map[int]int)
		// 무한 루프
		for {
			// reads 및 writes 채널 중 하나가 준비되면 해당 채널을 선택합니다.
			select {
				// reads 채널이 준비되면
				case read := <-reads:
					// read 채널에 readOp 구조체의 resp 필드에 state[read.key] 값을 전송
					read.resp <- state[read.key]
				// writes 채널이 준비되면
				case write := <-writes:
					// state[write.key]에 write.val 값을 할당
					state[write.key] = write.val
					// write.resp 채널에 true 값을 전송
					write.resp <- true
			}
		}
	}()

	// 100개의 고루틴을 생성하여 읽기 요청을 보냅니다.
	for r := 0; r < 100; r++ {
		go func() {
			// 무한 루프
			for {
				// readOp 구조체 생성
				read := readOp{
					// key는 0~4 사이의 랜덤한 정수
					key:  rand.Intn(5),
					// resp 채널 생성
					resp: make(chan int),
				}
				// reads 채널에 readOp 구조체 전송
				reads <- read
				// read 구조체의 resp 채널에서 값을 읽음
				// read.resp에 state[read.key] 값이 전송되기를 기다림
				<-read.resp
				// readOps 값을 1 증가
				atomic.AddUint64(&readOps, 1)
				// 1밀리초 대기
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10개의 고루틴을 생성하여 쓰기 요청을 보냅니다.
	for w := 0; w < 10; w++ {
		go func() {
			// 무한 루프
			for {
				// writeOp 구조체 생성
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				// writes 채널에 writeOp 구조체 전송
				writes <- write
				// writeOp 구조체의 resp 채널에서 값을 읽음
				// write.resp에 true 값이 전송되기를 기다림
				<-write.resp
				// writeOps 값을 1 증가
				atomic.AddUint64(&writeOps, 1)
				// 1밀리초 대기
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}