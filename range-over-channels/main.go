package main

// for와 range가 기본 데이터 구조에 대한 반복을 제공하는 방법을 살펴보았습니다.
// 이 구문을 사용하여 채널에서 받은 값을 반복할 수도 있습니다.
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 채널을 닫지 않으면 range 구문에서 deadlock 발생

	for elem := range queue {
		println(elem)
	}
}