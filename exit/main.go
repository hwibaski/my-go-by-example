package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exit() 호출은 프로그램을 즉시 종료합니다.
	// defer로 등록된 함수는 실행되지 않습니다.
	defer fmt.Println("exit")

	// C와 달리 Go는 종료 상태를 나타내기 위해 메인에서 정수 반환값을 사용하지 않습니다. 0이 아닌 상태로 종료하려면 os.Exit를 사용해야 합니다.
	os.Exit(3)
}
