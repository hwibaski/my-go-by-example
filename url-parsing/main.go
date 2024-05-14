package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=abc&k2=efg#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // postgres
	fmt.Println(u.User) // user:pass
	p, _ := u.User.Password()
	fmt.Println(p) // pass

	fmt.Println(u.Host) // host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host) // host.com
	fmt.Println(port) // 5432

	fmt.Println(u.Path) // /path
	// # 뒤에 오는 경로와 fragment 추출합니다.
	fmt.Println(u.Fragment) // f

	fmt.Println(u.RawQuery) // k=abc&k2=efg
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m) // map[k:[abc], k2:[eft]]
	fmt.Println(m["k"]) // [abc]
	fmt.Println(m["k"][0]) // abc
	// fmt.Println(m["k"][1]) // index out of range
	fmt.Println(m["k2"]) // [efg]
	fmt.Println(m["k2"][0]) // efg
}