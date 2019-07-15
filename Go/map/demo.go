package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/6/25 15:28
* @Name:_map
* @Function:
 */

func main() {
	mm := "qwer"
	m := map[int]*string{2: &mm}
	if v, ok := m[2]; ok {
		tt := "sdfsdf"
		v = &tt
		m[2] = &tt
		println(v)
	} else {

	}
	fmt.Printf("%v", *m[2])

	//
	mMap := make(map[string]interface{})
	mMap["23"] = 4
	mMap["234"] = "1234"

	fmt.Printf("%v", mMap["345345"])
}
