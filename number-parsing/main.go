package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string -> Float
	// 64는 구문 분석할 정밀도 비트를 나타냅니다.
	// https://pkg.go.dev/strconv#ParseFloat
	// 32를 전달해도 float64로 반환됩니다.
	f, _ := strconv.ParseFloat("1.234", 32)
	fmt.Println(f) // 1.234
	f32Var := float32(f)
	fmt.Println("f32Var:", f32Var) // 1.234

	// ParseInt의 경우 0은 문자열에서 기본값을 유추하는 것을 의미합니다.
	// 64는 결과가 64비트에 맞아야 합니다.
	// https://pkg.go.dev/strconv#ParseInt
	// 음수도 파싱 가능
	// base 인수가 0인 경우, 실제 기본값은 부호 뒤에 오는 문자열의 접두사(있는 경우)에 의해 암시됩니다: 
	// "0b"의 경우 2, "0" 또는 "0o"의 경우 8, "0x"의 경우 16, 그렇지 않은 경우 10입니다.
	// 두 번째 인수 : 주어진 문자열을 변환할 진수 (2에서 36 사이)입니다.
	// 세 번째 인수 : 결과가 표현되는 비트 크기 (0에서 64 사이)입니다.
	i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i) // 123

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d) // 456

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u) // 789

	// Atoi는 10진수 정수를 파싱합니다.
	k, _ := strconv.Atoi("135")
	fmt.Println(k) // 135

	// 잘못된 입력이 발생하면 Parse 함수는 오류를 반환합니다.
	_, e := strconv.Atoi("wat")
	fmt.Println(e) // strconv.Atoi: parsing "wat": invalid syntax
}