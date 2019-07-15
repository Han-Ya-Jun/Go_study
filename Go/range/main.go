package main

import (
	"fmt"
)

/*
* @Author:hanyajun
* @Date:2019/7/9 11:25
* @Name:_range
* @Function:
 */

func main() {
	type user struct {
		name string
	}

	s := make([]*user, 4, 4)
	s[0] = &user{"go"}
	s[1] = &user{"go"}
	s[2] = &user{"go"}
	s[3] = &user{"go"}

	// 先把v的name都修改为php
	for _, v := range s {
		v.name = "php"
	}
	fmt.Printf("%v", s[0])
}
