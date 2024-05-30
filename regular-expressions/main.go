package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 문자열을 바로 매칭
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	println(match) // true

	// 정규식 객체를 먼저 생성한 후 매정
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach")) // true
	fmt.Println(r.FindString("peach punch")) // peach
	// 일치하는 텍스트의 첫 인덱스와 마지막 인덱스를 반환합니다.
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // [0 5]

	// Submatch 변형은 전체 패턴 일치와 해당 일치의 부분 일치에 대한 정보를 모두 포함합니다.
	// 예를 들어 다음은 p([a-z]+)ch와 ([a-z]+)에 대한 정보를 모두 반환합니다
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// 전체 일치와 부분 일치의 인덱스에 대한 정보를 반환합니다.
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// All 변형은 입력에서 첫 번째 일치 항목만이 아닌 모든 일치 항목에 대한 정보를 반환합니다.
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// 이전과 유사하지만 각 일치 항목에 대한 인덱스 정보를 반환합니다.
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// 이 함수는 두 번째 인수로 양의 정수를 받아 해당 숫자만큼 일치 항목을 반환합니다.
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// 이 함수는 바이트 슬라이스를 받아서 일치 여부를 확인합니다.
	fmt.Println(r.Match([]byte("peach"))) // true

	// 정규표현식으로 상수를 만들때 Compile의 변형인 MustCompile을 사용할 수 있습니다.
	// 일반 Compile은 반환값이 2개라 상수로 사용할 수 없습니다.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp 패키지는 부분문자열을 다른값으로 바꾸는데 사용할 수도 있습니다.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	// Func 변형을 사용하여 주어진 함수로 일치된 텍스트를 변환시킬 수 있습니다.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}