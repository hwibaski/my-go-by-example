package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// rand.IntN은 임의의 정수 n, 0 <= n < 100을 반환합니다.
	fmt.Println(rand.IntN(100), ",")
	fmt.Println(rand.IntN(100))
	fmt.Println()
	
	// rand.Float64는 0.0 <= f < 1.0의 float64 f를 반환합니다.
	fmt.Println(rand.Float64())

	// 5.0 <= f' < 10.0과 같이 다른 범위에서 임의의 부동 소수점을 생성하는 데 사용할 수 있습니다.
	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 알려진 시드를 원한다면 새 rand.Source를 생성하고 이를 New 생성자에 전달하세요.
	// NewPCG는 두 개의 uint64 숫자로 된 시드가 필요한 새 PCG 소스를 생성합니다.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}