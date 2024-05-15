package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// /tmp/dat 파일 내용
// hello
// go

func main() {
	// 가장 기본적인 파일 읽기 작업은 파일의 전체 내용을 메모리에 저장하는 것입니다.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))
	// hello
	// go
	// 

	// 파일의 어떤 부분을 어떻게 읽을지 더 세밀하게 제어하고 싶을 때가 많습니다.
	// 이러한 작업의 경우 먼저 파일을 열어 os.File 값을 얻습니다.
	f, err := os.Open("/tmp/dat")
	check(err)

	// --------------------------------------------------------------------------------

	// 5byte 버퍼 생성
	b1 := make([]byte, 5)
	// 파일의 처음 5byte를 읽어서 버퍼에 저장
	// n1은 실제로 읽은 바이트 수를 나타냅니다.
	n1, err := f.Read(b1)
	check(err)
	fmt.Println(n1) // 5
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // 5 bytes: hello

	// --------------------------------------------------------------------------------
	/*
	 f.Seek()는 파일의 읽기/쓰기 위치를 변경하는 메서드입니다.
	 Go 언어에서 Seek 메서드는 io.Seeker 인터페이스에 정의되어 있으며, 파일이나 다른 데이터 스트림에서 현재 위치를 이동하는 데 사용됩니다.
	 func (f *File) Seek(offset int64, whence int) (int64, error)
     - offset: 이동하려는 위치. 이 값은 whence에 따라 해석됩니다.
	 - whence: 이동하려는 위치를 해석하는 방법. io.SeekStart, io.SeekCurrent, io.SeekEnd 중 하나입니다.
	   - io.SeekStart: 파일의 시작부터 offset만큼 이동합니다.
	   - io.SeekCurrent: 현재 위치에서 offset만큼 이동합니다.
	   - io.SeekEnd: 파일의 끝에서 offset만큼 이동합니다.
	*/

	// 파일의 처음부터 6바이트 떨어진 곳으로 이동
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	// 2byte 버퍼 생성
	b2 := make([]byte, 2)
	// 파일의 현재 위치에서 2byte를 읽어서 버퍼에 저장
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2])) // 2 bytes @ 6: go

	// --------------------------------------------------------------------------------

	// 파일의 현재 위치에서 4바이트 떨어진 곳으로 이동
	_, err = f.Seek(4, io.SeekCurrent)
    check(err)

	// 파일의 끝에서 10바이트 떨어진 곳으로 이끝
	_, err = f.Seek(-10, io.SeekEnd)
    check(err)

	// 파일의 처음부터 6바이트 떨어진 곳으로 이음
	o3, err := f.Seek(6, io.SeekStart)
    check(err)
	// --------------------------------------------------------------------------------
	// io.ReadAtLeast() 이용
	// ReadAtLast는 최소 min(3번째 인수) 바이트를 읽을 때까지 r(첫번째 인수)에서 buf(두 번째 인수)로 읽습니다.

	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3)) // 2 bytes @ 6: go


	_, err = f.Seek(0, io.SeekStart)
    check(err)

	// --------------------------------------------------------------------------------

	r4 := bufio.NewReader(f)
	// 주어진 리더에서 다음 n 바이트를 읽고 바이트 버퍼를 반환합니다.
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4)) // 5 bytes: hello

	// 작업을 마치면 파일을 닫습니다(일반적으로 이 작업은 defer를 사용하여 연 직후에 예약됩니다).
	f.Close()
}