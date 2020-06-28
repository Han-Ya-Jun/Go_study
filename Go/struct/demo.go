package main

import "fmt"

/*
* @Author: HanYaJun
* @Date: 2020/5/12 10:21 上午
* @Name: _struct
* @Function:
 */

type T struct {
	_ int
	_ bool
}

func main() {
	var t1 = T{123, true}
	var t2 = T{789, false}
	fmt.Printf("%v", t1)
	println(t1 == t2) // 打印出什么？
}
