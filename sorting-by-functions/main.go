package main

import (
	"cmp"
	"fmt"
	"slices"
)

// 컬렉션을 일반적인 순서가 아닌 다른 기준으로 정렬하고 싶을 때가 있습니다.
// 예를 들어 문자열을 알파벳순이 아닌 길이순으로 정렬하고 싶다고 가정해 보겠습니다.
// 다음은 Go에서 사용자 정의 정렬의 예입니다.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		// cmp.Compare를 사용하여 두 문자열의 길이를 비교합니다.
		// cmp.Compare는 두 인수를 비교하고 첫 번째 인수가 두 번째 인수보다 작으면 -1, 같으면 0, 크면 1을 반환합니다.
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age int
	}

	people := []Person{
		Person{"Jax", 37},
		Person{"TJ", 25},
		Person{"Alex", 72},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	// 참고: Person 구조가 큰 경우 슬라이스에 *Person을 대신 포함하도록 하고 그에 따라 정렬 기능을 조정할 수 있습니다.
	// 확실하지 않은 경우 벤치마킹하세요!
	fmt.Println(people)
}