package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")
	// 2024/05/28 01:22:58 standard logger

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")
	// 2024/05/28 01:22:58.096886 with micro

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")
	// 2024/05/28 01:22:58 main.go:18: with file/line

	mylog := log.New(os.Stdout, "mylog ", log.LstdFlags)
	mylog.Println("from mylog")
	// mylog 2024/05/28 01:22:58 from mylog

	mylog.SetPrefix("ohmy: ")
	mylog.Println("from mylog")
	// ohmy: 2024/05/28 01:22:58 from mylog

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf: ", log.LstdFlags)

	// 버퍼에 로그를 저장
	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())
	// from buflog:buf: 2024/05/28 01:24:50 hello

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("info")
	// {"time":"2024-05-28T01:24:50.819432+09:00","level":"INFO","msg":"info"}

	myslog.Info("hello again", "key", "val", "age", 25)
	// {"time":"2024-05-28T01:24:50.819442+09:00","level":"INFO","msg":"hello again","key":"val","age":25}
}
