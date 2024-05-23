package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// go test -v

// 일반적으로 테스트하는 코드는 intutils.go와 같은 이름의 소스 파일에 있으며,
// 이에 대한 테스트 파일은 intutils_test.go로 명명됩니다.
// Test로 시작하는 이름을 가진 함수를 작성하면 테스트가 만들어집니다.

func TestInMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error*는 테스트 실패를 보고하지만 테스트를 계속 실행합니다.
		// t.Fatal*은 테스트 실패를 보고하고 즉시 테스트를 중지합니다
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 테스트 작성은 반복적일 수 있으므로 테스트 입력과 예상 출력을 표에 나열하고
// 단일 루프가 그 위로 이동하여 테스트 로직을 수행하는 표 중심 스타일을 사용하는 것이 관용적입니다
func TestInMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.a, tt.b)
		// t.Run을 사용하면 각 테이블 항목에 대해 하나씩 "하위 테스트"를 실행할 수 있습니다.
		// 이는 go test -v를 실행할 때 별도로 표시됩니다.
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})

	}
}

// 벤치마크 테스트는 일반적으로 _test.go 파일에 저장되며 이름이 Benchmark로 시작됩니다.
// 테스트 러너는 각 벤치마크 함수를 여러 번 실행하여 정확한 측정값을 수집할 때까지 실행할 때마다 b.N을 증가시킵니다.

// go test -bench=.

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
