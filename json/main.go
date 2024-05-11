package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// to json
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) // true

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // "gopher"

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // ["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) // {"apple":5,"lettuce":7}

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}

	// --------------------------------------------
	// from json

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) // map[num:6.13 strs:[a b]]
	fmt.Printf("%T\n", dat["num"]) // float64

	// 디코딩된 맵의 값을 사용하려면 해당 값을 적절한 유형으로 변환해야 합니다. 
	// 예를 들어 여기에서는 num의 값을 예상되는 float64 유형으로 변환합니다.
	num := dat["num"].(float64)
	fmt.Println(num) // 6.13

	// 중첩된 데이터 변환
	strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1) // a
	
	// JSON을 사용자 정의 데이터 유형으로 디코딩할 수도 있습니다.
	// 이렇게 하면 프로그램에 타입 안전성을 추가하고 디코딩된 데이터에 액세스할 때 type assertion이 필요하지 않다는 이점이 있습니다.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)

    fmt.Println(res) // {1 [apple peach]}
    fmt.Println(res.Fruits[0]) // apple

	// 또한 JSON 인코딩을 os.Stdout과 같은 os.Writers나 심지어 HTTP 응답 본문으로 직접 스트리밍할 수도 있습니다.
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}