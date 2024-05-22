package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 임시 파일을 만드는 가장 쉬운 방법은 os.CreateTemp를 호출하는 것입니다. 이 함수는 파일을 생성하고 읽기 및 쓰기를 위해 파일을 엽니다.
	// 첫 번째 인수로 ""를 제공하면 os.CreateTemp가 OS의 기본 위치에 파일을 생성합니다.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// 임시 파일의 이름을 표시합니다.
	// 유닉스 기반 OS에서는 디렉터리가 /tmp일 가능성이 높습니다.
	// 파일 이름은 os.CreateTemp의 두 번째 인수로 지정된 접두사로 시작하고 나머지는 동시 호출이 항상 다른 파일 이름을 만들도록 자동으로 선택됩니다.
	fmt.Println("Temp file name:", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 많은 임시 파일을 작성하려는 경우 임시 디렉터리를 만드는 것이 좋습니다. 
	// os.MkdirTemp의 인수는 CreateTemp와 동일하지만 열린 파일이 아닌 디렉터리 이름을 반환합니다.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}