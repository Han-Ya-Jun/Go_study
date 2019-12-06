package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
	LoadWaybills []int `json:"load_waybills"`
}

const (
	test2 = 2
	test  = iota
	test3
)

func (i *IntString) UnmarshalJSON(value []byte) error {
	var result string
	if len(value) == 0 {
		return nil
	}

	if strings.Contains(fmt.Sprintf("%s", value), "\"") {
		result = string(value[1 : len(value)-1])
	} else {
		result = string(value)
	}
	m, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		return err
	}
	*i = IntString(m)
	return nil
}

func main() {
	println(test3)
	var a int = 90
	var result = &RequestParams{LoadWaybills: []int{a, 123}}
	b, _ := json.Marshal(result)
	fmt.Printf("marshal result:%s\n", b)
	var result2 RequestParam
	_ = json.Unmarshal(b, &result2)
	fmt.Printf("unmarshal result:%v", result2.LoadWaybills)
}
