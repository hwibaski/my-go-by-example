package main

import (
	"os"
	"os/exec"
	"syscall"
)

// 실행 중인 Go 프로세스에서 액세스할 수 있는 외부 프로세스가 필요할 때 이 작업을 수행합니다.
// 때로는 현재 Go 프로세스를 다른(아마도 Go가 아닌) 프로세스로 완전히 대체하고 싶을 때가 있습니다.
//
//	이를 위해 Go의 고전적인 exec 함수 구현을 사용하겠습니다.
func main() {
	// 이 예제에서는 ls를 실행하겠습니다.
	// Go에는 실행하려는 바이너리의 절대 경로가 필요하므로 exec.LookPath를 사용하여 해당 경로(아마도 /bin/ls)를 찾습니다.
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec에는 하나의 큰 문자열이 아닌 슬라이스 형식의 인수가 필요합니다.
	// 몇 가지 일반적인 인수를 ls에 제공하겠습니다. 첫 번째 인수는 프로그램 이름이어야 한다는 점에 유의하세요.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec은 또한 사용할 환경 변수 집합이 필요합니다. 여기서는 현재 환경만 제공합니다.
	env := os.Environ()

	// 다음은 실제 syscall.Exec 호출입니다.
	//  이 호출이 성공하면 프로세스 실행은 여기서 끝나고 /bin/ls -a -l -h 프로세스로 대체됩니다.
	//  에러가 발생하면 반환값이 반환됩니다.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
