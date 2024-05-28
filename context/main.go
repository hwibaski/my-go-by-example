package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// context.Context는 net/http 기계에 의해 각 요청에 대해 생성되며, Context() 메서드를 통해 사용할 수 있습니다.
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// 다음은 클라이언트에게 응답을 보내기 전에 몇 초간 대기하는 코드입니다.
	// 이 코드는 서버가 작업을 수행하는 것을 시뮬레이션할 수 있습니다.
	// 작업을 수행하는 동안, context의 Done() 채널에서 신호를 감시하여, 신호가 오면 작업을 취소하고 가능한 한 빨리 반환합니다.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// 컨텍스트의 Err() 메서드는 Done() 채널이 닫힌 이유를 설명하는 오류를 반환합니다.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

// &로 백그라운드에서 실행합니다.
// $ go run context-in-http-servers.go &

// $ curl localhost:8090/hello
// server: hello handler started
// ^C
// server: context canceled
// server: hello handler ended
