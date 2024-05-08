package main

import (
	"fmt"
	"slices"
)

func main() {
	// nil 슬라이스 생성
	var s []string

	fmt.Println("uninitialized slice:", s, s == nil, len(s) == 0) // uninitialized slice: [] true true

	// make 함수로 슬라이스 생성
	s = make([]string, 3)
	fmt.Println("empty slice:", s, "len:", len(s), "cap:", cap(s)) // empty slice: [  ] len: 3 cap: 3

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) // set: [a b c]
	fmt.Println("get:", s[2]) // get: c
	fmt.Println("len:", len(s)) // len: 3

	// append 함수로 슬라이스에 요소 추가
	// append 함수는 슬라이스에 새로운 요소를 추가하고 새로운 슬라이스를 반환한다.
	// append 함수에 슬라이스를 전달하면 헤당 슬라이스가 가리키는 배열에 요소를 추가한다.
	s = append(s, "d")
	f := append(s, "e", "f")
	fmt.Println("append:", s, "len:", len(s), "cap:", cap(s)) // append: [a b c d] len: 4 cap: 6
	fmt.Println("append:", f, "len:", len(f), "cap:", cap(f)) // append: [a b c d e f] len: 6 cap: 6

	// f 슬라이스와 같은 길이를 가지는 새로운 슬라이스를 생성
	c := make([]string, len(f))
	fmt.Println("c:", c, "len:", len(c), "cap:", cap(c)) // c: [     ] len: 6 cap: 6
	// copy 함수로 슬라이스 복사, c에 f를 복사
	copy(c, f)
	fmt.Println("copy:", c, "len:", len(c), "cap:", cap(c)) // copy: [a b c d e f] len: 6 cap: 6

	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s)) // slice: [a b c d] len: 4 cap: 6

	l := s[2:5]
	// slice s가 [a b c d] 여도 [c d e]가 반환된다. 그 이유는 내부적으로 가리키고 있는 배열이 [a b c d e f]이기 때문이다. f 슬라이스 확인!
	fmt.Println("sl1", l) // sl1 [c d e]

	l = s[:5]
	fmt.Println("sl2", l) // sl2 [a b c d e]

	l = s[2:]
	fmt.Println("sl3", l) // sl3 [c d]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // dcl: [g h i]

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal") 
	}
	// t and t2 are equal 출력, t1과 t2는 같은 요소를 가지고 있기 같은 슬라이스라고 판단.

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD) // 2d: [[0] [1 2] [2 3 4]]
} 