package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/5/14 9:49
* @Name:_switch
* @Function:
 */

func ff() bool {
	return false
}

func main() {
	// switch 默认为true
	switch {
	case true:
		fmt.Println("默认true")
	}

	// switch 支持初始化语句
	switch f := ff(); f {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("fasle")
	}

	// switch 只有初始化语句 条件默认为true
	switch ff(); {
	case true:
		fmt.Println("默认为true")
	case false:
		fmt.Println("默认为fasle")
	}

	// switch 只有初始化语句 条件默认为true 同上
	switch ff(); {
	case true:
		fmt.Println("默认为true")
	case false:
		fmt.Println("默认为fasle")
	}
}
