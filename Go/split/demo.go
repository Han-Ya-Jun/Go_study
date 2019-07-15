package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
* @Author:hanyajun
* @Date:2019/6/28 17:31
* @Name:split
* @Function:
 */

func main() {
	strs := strings.Split("1231-", "-")
	fmt.Println(len(strs), strs)
	fmt.Println(strs[1] == "")
	type Cc struct {
		Aa int `json:"aa,string"`
	}

	var cc = Cc{}
	var ccJson = `{"aa":"4"}`
	//var ccJson = `{"aa":"8"}`
	err := json.Unmarshal([]byte(ccJson), &cc)
	fmt.Println(err)
	fmt.Print(cc)
}
