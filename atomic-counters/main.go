package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Go에서 상태를 관리하는 주요 메커니즘은 채널을 통한 커뮤니케이션입니다.
// 하지만 상태 관리를 위한 몇 가지 다른 옵션이 있습니다.
// 여기서는 여러 고루틴에서 액세스하는 원자 카운터에 sync/atomic 패키지를 사용하는 방법을 살펴보겠습니다.

func main() {
	// atomic를 사용하여 고루틴 간 카운터를 관리합니다.
	var ops atomic.Uint64

	// 여러 고루틴을 기다기리 위한 WaitGroup 생성 
	var wg sync.WaitGroup

	// 50개의 고루틴을 생성하여 카운터를 증가시킵니다.
	for i := 0; i < 50; i++ {
		// wg.Add(1)을 사용하여 WaitGroup에 고루틴을 추가합니다.
		wg.Add(1)

		// 고루틴을 실행
		go func() {
			// 1000 번 실행
			for c := 0; c < 1000; c++ {
				// ops 카운터를 증가시킵니다.
				ops.Add(1)
			}
			// 고루틴이 끝나면 Done 메서드를 호출하여 WaitGroup에 알립니다.
			wg.Done()
		}()
	}

	// 모든 고루틴이 끝날 때까지 기다립니다.
	wg.Wait()

	// ops 카운터를 읽어옵니다.
	// atomic을 사용하지 않았다면 race condition이 발생할 수 있습니다.
	fmt.Println("ops:", ops.Load())
}