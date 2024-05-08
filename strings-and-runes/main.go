package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	// byte의 길이를 출력하기 때문에 18이 출력된다.
	// UTF-8은 영어 알파벳과 숫자는 1바이트, 한글은 3바이트, 이모티콘은 4바이트로 인코딩된다.
	fmt.Println("Len:", len(s)) // Len: 18

	// 문자열을 바이트 단위로 출력한다.
	for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i]) // e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
    }
    fmt.Println()

	// Rune은 UTF-8 문자열에서 하나의 문자를 나타내는 타입이다.
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6

	// range 키워드를 사용하여 rune을 출력한다.
	for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
	// U+0E2A 'ส' starts at 0
	// U+0E27 'ว' starts at 3
	// U+0E31 'ั' starts at 6
	// U+0E2A 'ส' starts at 9
	// U+0E14 'ด' starts at 12
	// U+0E35 'ี' starts at 15
	fmt.Println("\nUsing DecodeRuneInString")

	// DecodeRuneInString 함수를 사용하여 rune을 출력한다.
	for i, w := 0, 0; i < len(s); i += w {
		// DecodeRuneInString 함수는 문자열에서 첫 번째 문자를 읽어 rune과 rune의 길이를 반환한다.
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
	// U+0E2A 'ส' starts at 0
	// found so sua
	// U+0E27 'ว' starts at 3
	// U+0E31 'ั' starts at 6
	// U+0E2A 'ส' starts at 9
	// found so sua
	// U+0E14 'ด' starts at 12
	// U+0E35 'ี' starts at 15
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}