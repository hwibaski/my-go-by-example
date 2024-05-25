package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO")) // FOO: 1
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR:

	fmt.Println()
	// os.Environ은 환경 변수 목록을 키-값 쌍으로 반환합니다.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2) // 문자열을 구분자로 나누어 두 부분으로 나눕니다.
		fmt.Println(pair[0])              // 출력: FOO
	}
}
