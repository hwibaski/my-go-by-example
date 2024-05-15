package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println(data)
	// string -> []byte 변환 -> base64로 인코딩
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMTIzIT8kKiYoKSctPUB+

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // abc123!?$*&()'-=@~
	fmt.Println()

	// URL 호환 base64 형식을 사용하여 인코딩/디코딩합니다.
	// 문자열은 표준 및 URL base64 인코더(후행 + 대 -)를 사용하여 약간 다른 값으로 인코딩되지만 
	// 둘 다 원하는 대로 원래 문자열로 디코딩됩니다.
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc) // YWJjMTIzIT8kKiYoKSctPUB-

    uDec, _ := base64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec)) // abc123!?$*&()'-=@~
}