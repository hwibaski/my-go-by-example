package main

import (
	"fmt"
	"sync"
)

// 이전 예제에서는 원자 연산을 사용해 간단한 카운터 상태를 관리하는 방법을 살펴봤습니다.
// 더 복잡한 상태의 경우 뮤텍스를 사용하여 여러 고루틴에서 데이터에 안전하게 액세스할 수 있습니다.

// 컨테이너는 카운터 맵을 보유합니다.
// "여러 고루틴"에서 동시에 업데이트하고 싶기 때문에 액세스를 동기화하기 위해 Mutex를 추가합니다.
// 참고 : 뮤텍스는 복사할 수 없으므로 이 구조체를 전달할 때는 포인터로 전달해야 합니다.
type Container struct {
	mu 		 sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// mutex를 사용하여 연산을 하기전 뮤텍스를 잠급니다.
	c.mu.Lock()
	// defer를 사용하여 뮤텍스를 해제합니다.
	defer c.mu.Unlock()
	// 카운터를 증가시킵니다.
	c.counters[name]++
}

func main() {
	// 컨테이너를 생성하고 카운터를 초기화합니다.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	// 여러 고루틴을 컨트롤 하기 위해 WaitGroup을 생성합니다.
	var wg sync.WaitGroup

	// Container의 inc 메서드를 호출하는 함수를 생성
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// 3개의 고루틴을 생성하여 카운터를 증가시킵니다.
	wg.Add(3)

	// a 카운터를 10000번 증가시키는 고루틴을 생성합니다.
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	// b 카운터를 10000번 증가시키는 고루틴을 생성합니다.
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters) // map[a:20000 b:10000]
}