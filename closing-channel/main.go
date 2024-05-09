package main

import "fmt"

// 채널을 닫는 것은 해당 채널로 더 이상 값을 보내지 않겠다는 것을 나타냅니다.
// 이는 채널의 수신자에게 완료를 알리는 데 유용할 수 있습니다.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 고루틴으로 무한 루프를 돌면서 채널로부터 값을 받아옵니다.
	go func() {
		for {
			j, more := <-jobs
			// 채널이 닫히면 more는 false가 됩니다.
			if more {
				// 채널의 값 출력
				fmt.Println("received job", j)
			} else {
				// 채널이 닫혔을 경우
				fmt.Println("received all jobs")
				// done 채널에 true를 보냅니다.
				done <- true
				// 무한 루프를 종료합니다.
				return
			}
		}
	}()

	// jobs 채널로 3개의 값을 보냅니다.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// jobs 채널을 닫습니다.
	close(jobs)
	fmt.Println("sent all jobs")

	// done 채널에서 값을 받을 때 까지 대기합니다.
	// 동기화를 위해 사용되었습니다.
	<-done

	// 채널이 닫혔는지 확인합니다.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}