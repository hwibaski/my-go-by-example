package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// line filter는 stdin에서 입력을 읽고 처리한 다음 파생된 결과를 stdout에 출력하는 일반적인 유형의 프로그램입니다.
// grep과 sed가 일반적인 line filter입니다.
// 다음은 모든 입력 텍스트의 대문자 버전을 작성하는 Go의 line filter 예제입니다.
// 이 패턴을 사용하여 자신만의 Go line filter를 작성할 수 있습니다.

// echo "hello\nfilter" | go run line-filters/main.go
func main() {
	// 버퍼되지 않은 os.Stdin을 버퍼링된 스캐너로 감싸면
	// 스캐너가 기본 스캐너의 다음 줄인 다음 토큰으로 이동하는 편리한 스캔 메서드를 사용할 수 있습니다.
	scanner := bufio.NewScanner(os.Stdin)

	// 입력에서 현재 토큰(여기서는 다음 줄)을 반환합니다.
	for scanner.Scan() {
		// scanner.Text() : scanner.Scan()에서 걸린 줄을 반환한다
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(scanner.Text()) // hello
		fmt.Println(ucl) // Hello
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}