package main

import (
	"fmt"
)

/*
* @Author:hanyajun
* @Date:2019/6/25 15:28
* @Name:_map
* @Function:
 */

func main() {

	//
	mMap := make(map[string]interface{})
	mMap["23"] = 4
	mMap["234"] = "1234"
	if value, ok := mMap["23"]; ok {
		value = "5"
		fmt.Println(value)
	} else {
		fmt.Println(value)
	}
	fmt.Printf("%v", mMap["23"])
}
