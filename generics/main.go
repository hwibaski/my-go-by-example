package main

import "fmt"

// K, V는 타입 매개변수
// K는 비교 가능한 타입이어야 하고, V는 어떤 타입이든 상관 없다.
func MapKeys[K comparable, V any](m map[K]V) []K {
	// m의 키를 저장할 슬라이스를 생성
	r := make([]K, 0, len(m))

	// m의 키를 슬라이스에 추가
	for k := range m {
		r = append(r, k)
	}

	// 슬라이스 반환
	return r
}

// T는 임의의 타입
type List[T any] struct {
	// head와 tail은 element[T] 타입의 포인터
	head, tail *element[T]
}

// element 구조체
// T는 임의의 타입
type element[T any] struct {
	// next는 element[T] 타입의 포인터
	next *element[T]
	// val은 T 타입
	val  T
}

func (list *List[T]) Push(val T) {
	// list.tail 이 없으면
	if list.tail == nil {
		// list.head와 list.tail에 새 element를 할당
		list.head = &element[T]{val : val}
		list.tail = list.head
	} else {
		// list.tail.next에 새 element를 할당하고 list.tail을 새 element로 변경
		list.tail.next = &element[T]{val : val}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	// nil 슬라이스 생성
	var elems []T

	// list.head부터 시작해서 끝까지 반복
	for e := list.head; e != nil; e = e.next {
		// 슬라이스에 element 추가
		elems = append(elems, e.val)
	}

	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	// generic 타입 변수를 이용하여 타입 검사
	_ = MapKeys[int, string](m)

	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	fmt.Println("list:", list.GetAll())
}

// keys: [1 2 4]
// list: [10 13 23]