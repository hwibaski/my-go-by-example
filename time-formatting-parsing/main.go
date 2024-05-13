package main

import (
	"fmt"
	"time"
)

// Go는 패턴 기반 레이아웃을 통해 시간 서식 지정 및 구문 분석을 지원합니다.

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339)) // 2021-08-10T15:52:00+09:00

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	p(t1) // 2012-11-01 22:08:41 +0000 +0000
	p(t.Format("3:04PM")) // 5:34PM (현재 시간)
	p(t.Format("Mon Jan _2 15:04:05 2006")) // Mon May 13 17:34:09 2024 
    p(t.Format("2006-01-02T15:04:05.999999-07:00")) // 2024-05-13T17:34:09.871892+09:00
	form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2) // 0000-01-01 20:41:00 +0000 UTC

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", 
	t.Year(), t.Month(), t.Day(),
	t.Hour(), t.Minute(), t.Second())  // 2024-05-13T17:34:09-00:00

	ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e) // parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"

}