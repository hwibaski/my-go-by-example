package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// 센티넬 에러 : 에러를 변수로 정의하여 사용하는 방법
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")
var TempErr = fmt.Errorf("making tea: %w", ErrPower)

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// %w 포맷을 사용하여 에러를 래핑한다.
		// var TempErr = fmt.Errorf("making tea: %w", ErrPower)
		return fmt.Errorf("temp: %w", TempErr)
	}

	return nil
}

func main() { 
	for _, i := range []int{7, 42} {

		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// errors.Is 함수를 사용하여 에러를 비교한다.
			// errors.Is 함수는 래핑된 에러도 재귀적으로 비교할 수 있다.
			// errors.Is는 주어진 오류(또는 해당 오류 체인의 모든 오류)가 특정 오류 값과 일치하는지 확인합니다.
			// 이 기능은 래핑되거나 중첩된 오류에 특히 유용하며, 오류 체인에서 특정 오류 유형이나 센티널 오류를 식별할 수 있습니다.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("i:", i ,"We should by new tea!")
			} else if errors.Is(err, TempErr) {
				fmt.Println(err)
				fmt.Println("i:", i, "Now it is dark.")
			} else {
				fmt.Printf("i:", i, "Unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}