package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	// jobs 채널 생성
	jobs := make(chan int, numJobs)
	// results 채널 생성, 작업한 결과를 받는 채널
	results := make(chan int, numJobs)

	// 3개의 워커를 생성합니다.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5개의 작업을 jobs 채널에 보냅니다.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// 모든 작업을 보냈음을 알리기 위해 jobs 채널을 닫습니다.
	close(jobs)

	// 각 작업의 결과를 수집합니다.
	for a := 1; a <= numJobs; a++ {
		// results 채널에서 값을 받습니다.
		// <-results는 결과를 받을 때까지 블록됩니다.
		<-results
	}
	fmt.Println("All jobs are done")
}