package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
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

	var x uint64 = 897797896767
	var y int64 = int64(x)
	fmt.Println(y)
	//fmt.Printf("uint64: %v = %#[1]x, int64: %v = %#x\n", x, y, uint64(y))

}
