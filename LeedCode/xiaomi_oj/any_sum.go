package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
* @Author:hanyajun
* @Date:18-11-9 下午5:15
* @Name:Go_study/main
* @Function:找出可能的合的组合
 */

func solution(line string) string {
	// 在此处理单行数据

	// 返回处理后的结果
	return ""
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
