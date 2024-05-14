package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 해시는 바이너리 또는 텍스트 블롭의 짧은 ID를 계산하는 데 자주 사용됩니다.
// 예를 들어, TLS/SSL 인증서는 SHA256을 사용하여 인증서의 서명을 계산합니다.
// Go에서 SHA256 해시를 계산하는 방법은 다음과 같습니다.

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	// 슬라이스에 추가하는 데 사용할 수 있지만 일반적으로는 필요하지 않습니다.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs) // 바이트 슬라이스로 출력
	fmt.Printf("%x\n", bs) // 16진수로 출력
}