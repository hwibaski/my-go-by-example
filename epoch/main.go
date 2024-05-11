package main

import (
	"fmt"
	"time"
)

// 프로그램에서 흔한 요구 사항 중 하나는 Unix epoch 이후의 초, 밀리초 또는 나노초 수를 가져오는 것입니다.
// Go에서는 다음과 같이 수행할 수 있습니다.

func main() {
	now := time.Now()
	fmt.Println("now:", now)

	// time.Now을 Unix, UnixMilli 또는 UnixNano와 함께 사용하여 Unix epoch 이후의 경과 시간을 각각 초, 밀리초 또는 나노초로 얻을 수 있습니다.
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// unix epoch 이후의 초, 밀리초 또는 나노초를 사용하여 time을 만들 수도 있습니다.
	fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}