package main

import (
	"errors"
	"fmt"
)


// 사용자 지정 에러는 주로 접미사에 Error를 붙여서 정의한다.
type argError struct {
	arg 	int
	message string
}

// error 인터페이스를 구현해야 한다.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	// errors.As는 errors.Is의 고급 버전입니다.
	// 주어진 오류(또는 그 체인의 모든 오류)가 특정 오류 유형과 일치하는지 확인하고 해당 유형의 값으로 변환하여 참을 반환합니다.
	// 일치하는 항목이 없으면 false를 반환합니다.	

	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err does not match argError")
	}
}