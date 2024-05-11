package main

import (
	"os"
	"text/template"
)

// Go는 text/template 패키지를 통해 동적 콘텐츠를 만들거나 사용자에게 맞춤화된 출력을 표시할 수 있는 기본 지원을 제공합니다.
// html/template이라는 이름의 형제 패키지는 동일한 API를 제공하지만 추가 보안 기능이 있으며 HTML을 생성하는 데 사용해야 합니다.

func main() {
	// 새 템플릿을 만들고 문자열에서 body를 파싱할 수 있습니다.
	// 템플릿은 정적 텍스트와 {{...}}로 묶인 "작업"이 혼합된 것으로, 이를 사용하여 콘텐츠를 동적으로 삽입합니다.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 또는 template.Must 함수를 사용하여 Parse가 오류를 반환하는 경우 패닉을 발생시킵니다.
	// 이 함수는 전역 범위에서 초기화된 템플릿에 특히 유용합니다.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 템플릿을 Execute 하면 해당 작업에 대한 특정 값이 포함된 텍스트가 생성됩니다. 
	// {{.}} 액션은 Execute에 매개변수로 전달된 값으로 대체됩니다.
	t1.Execute(os.Stdout, "some text") // Value: some text 출력됨
	t1.Execute(os.Stdout, 5) // Value: 5 출력됨
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"}) // Value: [Go Rust C++ C#] 출력됨

	// --------------------------------------------------------------------------------

	// 템플릿의 이름과 템플릿 문자열을 전달하여 새 템플릿을 만드는 함수를 만들 수 있습니다
	Create := func(name, t string) *template.Template {
		// 템플릿을 만들고 반환합니다.
		// Parse가 오류를 반환하는 경우 패닉을 발생시킵니다.
        return template.Must(template.New(name).Parse(t))
    }

	// 데이터가 구조체인 경우 {{.FieldName}} 액션을 사용하여 해당 필드에 액세스할 수 있습니다
	//  템플릿이 실행 중일 때 액세스할 수 있도록 필드를 내보내야 합니다.
    t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {Name string}{"Jane Doe"})

    t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse",})

	// --------------------------------------------------------------------------------

	// f/else는 템플릿에 조건부 실행을 제공합니다.
	// 값이 0, 빈 문자열, nil 포인터 등과 같은 유형의 기본값인 경우 거짓으로 간주됩니다.
	// 이 샘플은 템플릿의 또 다른 기능인 -를 사용하여 공백을 다듬는 동작을 보여줍니다.

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty") // yes
	t3.Execute(os.Stdout, "") // no

	// --------------------------------------------------------------------------------

	// range 블록을 사용하면 슬라이스, 배열, 맵 또는 채널을 반복할 수 있습니다.
	// range 블록 내부의 {{.}}는 반복의 현재 항목으로 설정됩니다.
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}