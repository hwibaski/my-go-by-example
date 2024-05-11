package main

import (
	"fmt"
	"os"
)

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}

	// %v : 기본 형식 출력 (struct1: {1 2})
    fmt.Printf("struct1: %v\n", p)

	// %+v : 구조체 필드 이름과 값을 출력 (struct2: {x:1 y:2})
    fmt.Printf("struct2: %+v\n", p)

	// %#v : 해당 값을 생성하는 소스 코드 스니펫 출력 (struct3: main.point{x:1, y:2})
    fmt.Printf("struct3: %#v\n", p)

	// %T : 해당 값의 타입 출력 (type: main.point)
    fmt.Printf("type: %T\n", p)

	// 
    fmt.Printf("bool: %t\n", true)

	// %d : 10진수 정수 출력 (int: 123)
    fmt.Printf("int: %d\n", 123)

	// %b : 2진수 출력 (bin: 1110)
    fmt.Printf("bin: %b\n", 14)

	// %c : 해당 유니코드 코드 포인트에 대응하는 문자 출력 (char: !)
    fmt.Printf("char: %c\n", 33)

	// %x : 16진수 출력 (hex: 1c8)
    fmt.Printf("hex: %x\n", 456)

	// %f : 기본적인 실수 출력 (float1: 78.900000)
    fmt.Printf("float1: %f\n", 78.9)

	// %e : 지수 표현 출력 (float2: 1.234000e+08)
    fmt.Printf("float2: %e\n", 123400000.0)

	// %E : 지수 표현 출력 (float3: 1.234000E+08)
    fmt.Printf("float3: %E\n", 123400000.0)

	// %s : 기본 문자열 출력 (str1: "string")
    fmt.Printf("str1: %s\n", "\"string\"")

	// %q : Go 구문을 따르는 문자열 출력 (str2: "\"string\"")	
    fmt.Printf("str2: %q\n", "\"string\"")

	// %x : 16진수 출력 (str3: 6865782074686973)
	//  입력 바이트당 두 개의 출력 문자를 사용하여 기본 16진수로 문자열을 렌더링합니다.
    fmt.Printf("str3: %x\n", "hex this")

	// %p : 포인터 출력 (pointer: 0xc0000140a0)
    fmt.Printf("pointer: %p\n", &p)

	// %6d : 최소 너비 6으로 오른쪽 정렬된 정수 출력 (width1: |    12|   345|)
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// %6.2f : 최소 너비 6, 소수점 이하 2자리로 오른쪽 정렬된 실수 출력 (width2: |  1.20|  3.45|)
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// %-6.2f : 최소 너비 6, 소수점 이하 2자리로 왼쪽 정렬된 실수 출력 (width3: |1.20  |3.45  |)
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// %6s : 최소 너비 6으로 오른쪽 정렬된 문자열 출력 (width4: |   foo|     b|)
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// %-6s : 최소 너비 6으로 왼쪽 정렬된 문자열 출력 (width5: |foo   |b     |)
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// fmt.Sprintf : 문자열 포맷팅 (sprintf: a string)
    s := fmt.Sprintf("sprintf: a %s", "string")
	// sprintf: a string
    fmt.Println(s)

	// Fprintf를 사용하여 os.Stdout이 아닌 io.Writers로 포맷+인쇄할 수 있습니다.
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}