package main

import (
	"flag"
	"fmt"
	"os"
)

// go 도구나 git과 같은 일부 명령줄 도구에는 각각 고유한 플래그 집합을 가진 많은 하위 명령이 있습니다.
// 예를 들어, go build와 go get은 go 도구의 서로 다른 두 가지 하위 명령어입니다.
// flag 패키지를 사용하면 자체 플래그가 있는 간단한 하위 명령을 쉽게 정의할 수 있습니다.

func main() {
	// NewFlagSet 함수를 사용하여 하위 명령을 선언하고 이 하위 명령에 특정한 새 플래그를 정의합니다.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// 하위 명령은 프로그램의 첫 번째 인수로 예상됩니다.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// $ ./command-line-subcommands foo -enable -name=joe a1 a2
// subcommand 'foo'
// enable: true
// name: joe
// tail: [a1 a2]

// $ ./command-line-subcommands bar -level 8 a1
// subcommand 'bar'
// level: 8
// tail: [a1]

// $ ./command-line-subcommands bar -enable a1
// flag provided but not defined: -enable
// Usage of bar:
// -level int
// level
