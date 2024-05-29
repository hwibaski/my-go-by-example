package main

import (
	"fmt"
	"io"
	"os/exec"
)

// Sometimes our Go programs need to spawn other, non-Go processes.
func main() {
	// 터미널에 date 명령을 실행하기 위함
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	println("> date")
	println(string(dateOut))

	// Command의 Output 및 기타 메서드는 명령 실행에 문제가 있는 경우(예: 잘못된 경로) *exec.Error를 반환하고,
	// 명령이 실행되었지만 반환 코드가 0이 아닌 상태로 종료된 경우 *exec.ExitError를 반환합니다.
	_, err = exec.Command("ls", "-a").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// grep 명령어를 실행할 준비를 합니다. grep 명령어는 "hello" 문자열을 찾는 역할을 합니다.
	grepCmd := exec.Command("grep", "hello")

	// grepCmd.StdinPipe()와 grepCmd.StdoutPipe()를 호출하여 grep 프로세스의 표준 입력(stdin)과 표준 출력(stdout)을 파이프합니다.
	//// 이를 통해 프로그램은 grep 명령어와 데이터를 주고받을 수 있습니다.
	//grepIn은 표준 입력 파이프를 나타내고, grepOut은 표준 출력 파이프를 나타냅니다.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	// 이 줄은 grep 프로세스를 시작합니다. 이제 grep 프로세스는 실행 중입니다.
	grepCmd.Start()

	// 이 줄은 grep 프로세스의 표준 입력으로 문자열 "hello grep\ngoodbye grep"을 씁니다.
	// grep 프로세스가 이 데이터를 입력으로 받아 처리하게 됩니다.
	grepIn.Write([]byte("hello grep\ngoodbye grep"))

	// 표준 입력 파이프를 닫습니다. 이는 더 이상 입력할 데이터가 없음을 grep 프로세스에 알립니다.
	grepIn.Close()

	// grep 프로세스의 표준 출력에서 모든 데이터를 읽어옵니다.
	// grep 프로세스는 입력 데이터에서 "hello"가 포함된 줄을 출력할 것이므로, 이 출력 데이터를 grepBytes에 저장합니다.
	grepBytes, _ := io.ReadAll(grepOut)

	// grep 프로세스가 종료될 때까지 기다립니다. 이 명령어는 grep 프로세스가 모든 작업을 마치고 종료될 때까지 블로킹됩니다.
	grepCmd.Wait()

	println("> grep hello")
	println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

	// 파이프를 사용하기 위한 방법
	// https://jusths.tistory.com/196
}
