package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// WaitGroup은 고루틴이 모두 끝날 때까지 기다리는 동안 대기할 수 있는 도구입니다.
	// 여러 개의 고루틴이 모두 끝날 때까지 기다리기 위해 사용됩니다
	// 참고: WaitGroup이 함수에 명시적으로 전달되는 경우, 포인터로 전달되어야 합니다.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// WaitGroup에 고루틴이 하나 추가됨을 알립니다.
		wg.Add(1)

		go func() {
			// 고루틴이 끝나면 Done 메서드를 호출하여 WaitGroup에 알립니다.
			defer wg.Done()
			// 워커 호출
			worker(i)
		}()
	}

	// 대기 그룹 카운터가 0으로 돌아갈 때까지 차단하고 모든 작업자에게 완료 알림을 보냅니다.
	// 이 접근 방식은 워커로부터 오류를 전파하는 직접적인 방법이 없다는 점에 유의하세요.
	// 보다 고급 사용 사례의 경우 errgroup 패키지를 사용하는 것이 좋습니다.
	wg.Wait()
}