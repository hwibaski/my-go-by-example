package main

import (
	"flag"
	"fmt"
)

// 명령줄 플래그는 명령줄 프로그램에 대한 옵션을 지정하는 일반적인 방법입니다.
//예를 들어, wc -l에서 -l은 명령줄 플래그입니다.

func main() {
	// flag의 이름, 기본값, 설명을 사용하여 새로운 플래그를 선언합니다.
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// 프로그램의 다른 곳에서 선언된 기존 변수를 사용하는 옵션을 선언할 수도 있습니다.
	// 플래그 선언 함수에 대한 포인터를 전달해야 한다는 점에 유의하세요.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Parse()를 호출하여 명령줄을 분석합니다.
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// go build command-line-flags.go

// $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
// word: opt
// numb: 7
// fork: true
// svar: flag
// tail: []

// $ ./command-line-flags -word=opt
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: []

// $ ./command-line-flags -word=opt a1 a2 a3
// word: opt
// ...
// tail: [a1 a2 a3]

// ********* 후행 인수 (a1 a2 a3) 앞에 flag 인수가 위치해야함 그렇지 않으면 제대로 파싱되지 않음 *********
// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

// $ ./command-line-flags -h
// Usage of ./command-line-flags:
// -fork=false: a bool
// -numb=42: an int
// -svar="bar": a string var
// -word="foo": a string

// 플래그 패키지에 지정되지 않은 플래그를 제공하면 프로그램에서 오류 메시지를 인쇄하고 도움말 텍스트를 다시 표시합니다.
// $ ./command-line-flags -wat
// flag provided but not defined: -wat
// Usage of ./command-line-flags:
// ...
