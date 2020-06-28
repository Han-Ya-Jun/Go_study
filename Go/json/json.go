package main

import (
	"encoding/json"
	"fmt"
)

/*
* @Author:hanyajun
* @Date:2019/8/6 10:21
* @Name:json
* @Function:
 */

type IntString int64

type RequestParam struct {
	LoadWaybills []IntString `json:"load_waybills"`
}
type RequestParams struct {
	LoadWaybills []int64 `json:"load_waybills"`
}

const (
	test2 = 2
	test  = iota
	test3
)

func main() {
	println(test3)
	var a int64 = 72059299143909760
	var result = &RequestParams{LoadWaybills: []int64{a, 123}}
	b, _ := json.Marshal(result)
	fmt.Printf("marshal result:%s\n", b)
	var result2 RequestParam
	_ = json.Unmarshal(b, &result2)
	fmt.Printf("unmarshal result:%v", result2.LoadWaybills)
}
