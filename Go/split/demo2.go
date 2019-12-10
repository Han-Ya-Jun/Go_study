package main

import (
	"fmt"
	"strings"
)

/*
* @Author:hanyajun
* @Date:2019/7/30 23:42
* @Name:main
* @Function:
 */

func main() {
	str := "a   b    c  de"
	sList := strings.Fields(str, " ")
	fmt.Println(sList)
}
