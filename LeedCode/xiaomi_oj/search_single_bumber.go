package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution1(line string) string {
	// 在此处理单行数据
	var term, value int
	strArray := strings.Split(line, " ")
	for _, v := range strArray {
		value, _ = strconv.Atoi(v)
		term = term ^ value
	}
	// 返回处理后的结果
	return fmt.Sprintf("%d", term)
}

func getSingleNumber(array []int) int {
	var term int
	for _, v := range array {
		term = term ^ v
	}
	return term
}

func main() {

	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution1(string(line)))
	}
}
