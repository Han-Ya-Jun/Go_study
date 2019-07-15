package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

/*
* @Author:hanyajun
* @Date:2019/6/28 15:51
* @Name:id
* @Function: id生成器
 */
func main() {
	n, err := snowflake.NewNode(2)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()

		fmt.Println("id", id)
	}
}
