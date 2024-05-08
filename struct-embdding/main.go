package main

import "fmt"

type base struct {
	num int
}

func (b base)  describe() string {
	// %v : 값을 기본 형식으로 표시, 정수는 10진수, 문자열은 문자열 그대로
	// %d : 정수를 10진수로 표시
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base // base 구조체를 임베딩
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// co는 container 타입이지만 base 타입도 포함하고 있기 때문에 describer 인터페이스를 구현한다.
	var d describer = co
	fmt.Println("describer:", d.describe())
}

// co={num: 1, str: some name}
// also num: 1
// describe: base with num=1
// describer: base with num=1