package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/8/8 17:38
* @Name:methond_receiver
* @Function:
 */

type MyStruct struct {
	x int
}

func (m MyStruct) Set1() {
	m.x = 1
}

func (m *MyStruct) Set2() {
	m.x = 2
}
func main() {

	s := &MyStruct{x: 0}
	s.Set1()
	fmt.Println(s.x)
	s.Set2()
	fmt.Println(s.x)
}
