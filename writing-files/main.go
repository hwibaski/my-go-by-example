package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 바이트 슬라이스 생성
	d1 := []byte("hello\ngo\n")

	// 기본적으로 os.WriteFile을 이용해서 파일에 씁니다.
	err := os.WriteFile("./dat1", d1, 0644)
	check(err)

	// 파일을 열어서 파일에 쓰는 작업을 더 세밀하게 제어하려면 파일을 열어 os.File 값을 얻어야 합니다.
	f, err := os.Create("./dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10} // "some\n"
	// 파일에 바이트 슬라이스를 씁니다.
	n2, err := f.Write(d2)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString도 사용할 수 있습니다.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// 동기화를 실행하여 안정적인 스토리지로 쓰기를 플러시합니다.
	f.Sync()

	// bufio는 앞서 살펴본 버퍼링된 리더 외에 버퍼링된 라이터도 제공합니다.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// 플러시를 사용하여 모든 버퍼링된 작업이 기본 라이터에 적용되었는지 확인합니다.
	w.Flush()
}