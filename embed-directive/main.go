package main

import "embed"

// //go:embed는 컴파일러 지시어로, 프로그램이 빌드 시 Go 바이너리에 임의의 파일과 폴더를 포함할 수 있게 해줍니다.
// embed 지시어에 대한 자세한 내용은 여기를 참조하세요.
// https://pkg.go.dev/embed

// embed 지시어는 go 소스 파일에 포함된 디렉터리 기준 상대 경로로 파일을 포함할 수 있습니다.
//
//go:embed folder/single_file.txt
var fileString string

// 또는 파일의 바이트 슬라이스를 포함할 수 있습니다.
//
//go:embed folder/single_file.txt
var fileByte []byte

// 여러 파일이나 폴더를 와일드카드를 사용하여 임베드할 수도 있습니다.
// 이를 위해 embed.FS 타입의 변수를 사용하는데, 이 타입은 간단한 가상 파일 시스템을 구현합니다.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
