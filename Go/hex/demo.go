package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/7/15 18:23
* @Name:hex
* @Function:
 */

func main() {
	//var a int64
	s := fmt.Sprintf("%0x", 9223372036854775807)
	fmt.Println(s)
}
