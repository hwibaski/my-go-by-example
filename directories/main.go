package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 디렉토리를 만듭니다. 0755는 디렉토리에 대한 기본 권한입니다.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// os.RemoveAll은 전체 디렉터리 트리를 삭제합니다(rm -rf와 유사하게).
	defer os.RemoveAll("subdir")

	// 빈 파일을 만들어주는 함수
	createdEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	// 빈 파일을 만듭니다.
	createdEmptyFile("subdir/file1")
	
	// 계층적 디렉토리 생성
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createdEmptyFile("subdir/parent/file2")
	createdEmptyFile("subdir/parent/file3")
	createdEmptyFile("subdir/parent/child/file4")

	// os.ReadDir은 디렉토리 내용을 읽습니다.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 현재 디렉터리의 위치를 변경
	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
    check(err)

	fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	err = os.Chdir("../../..")
    check(err)

	fmt.Println("Visiting subdir")
	// 또한 모든 하위 디렉터리를 포함하여 디렉터리를 재귀적으로 방문할 수도 있습니다.
	// WalkDir은 방문한 모든 파일 또는 디렉토리를 처리하는 콜백 함수를 허용합니다.
	err = filepath.WalkDir("subdir", visit)
}

func visit(p string, b fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, b.IsDir())
	return nil
}