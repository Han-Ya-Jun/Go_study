package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
* @Author:hanyajun
* @Date:2018/10/9 9:07
* @Name:go-study/main
* @Function:动态规划解决交叉队列问题
 */

func solution(line string) string {
	// 在此处理单行数据
	strArray := strings.Split(line, ",")
	if len(strArray[0])+len(strArray[1]) != len(strArray[2]) {
		return fmt.Sprintf("%v", false)
	}
	if len(strArray[0]) == 0 {
		return fmt.Sprintf("%v", strArray[1] == strArray[2])
	}
	if len(strArray[1]) == 0 {
		return fmt.Sprintf("%v", strArray[0] == strArray[2])
	}
	var intArray [len(strArray[0]) + 1][len(strArray[2]) + 1]int

	intArray[0][0] = 1
	// 返回处理后的结果
	return
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
