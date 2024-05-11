package main

import "fmt"

// Go에서는 내장된 recover 기능을 사용하여 패닉 상태에서 복구할 수 있습니다.
// recover 기능을 사용하면 패닉으로 인한 프로그램 중단을 중지하고 대신 실행을 계속할 수 있습니다.

// 클라이언트 연결 중 하나에 심각한 오류가 발생하는 경우 서버가 충돌을 일으키고 싶지 않을 때를 예로 들 수 있습니다.
// 대신 서버는 해당 연결을 닫고 다른 클라이언트를 계속 서비스하고 싶을 것입니다.
// 실제로 Go의 net/http는 HTTP 서버에 대해 기본적으로 이 작업을 수행합니다.

func mayPanic() {
	panic("a problem")
}

func main() {
	// recover는 반드시 defer 함수 내에서 호출되어야 합니다. 
	// 둘러싸는 함수가 패닉에 빠지면 디퍼가 활성화되고 그 안에 있는 recover 호출이 패닉을 잡습니다.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}