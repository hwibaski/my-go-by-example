package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p) // p: dir1/dir2/filename

	// 수동으로 /s 또는 \s를 연결하는 대신 항상 Join을 사용해야 합니다.
	// Join은 이식성을 제공할 뿐만 아니라 불필요한 구분 기호 및 디렉터리 변경을 제거하여 경로를 정규화합니다.
	fmt.Println(filepath.Join("dir1//", "filename")) // dir1/filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1/filename

	fmt.Println("Dir(p):", filepath.Dir(p)) // Dir(p): dir1/dir2
	fmt.Println("Base(p):", filepath.Base(p)) // Base(p): filename

	fmt.Println(filepath.IsAbs("dir/file")) // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext) // .json

	fmt.Println(strings.TrimSuffix(filename, ext)) // config

	// Rel은 기준과 대상 사이의 상대 경로를 찾습니다. 
	// 타깃이 기준과 상대적일 수 없는 경우 오류를 반환합니다.

	// base 에서 target 으로 가는 상대 경로를 반환
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // t/file

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // ../c/t/file

}