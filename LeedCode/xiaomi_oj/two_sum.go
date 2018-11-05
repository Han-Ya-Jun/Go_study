package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution0(line string) string {
	// 在此处理单行数据
	strArray := strings.Split(line, " ")
	a, _ := strconv.Atoi(strArray[0])
	b, _ := strconv.Atoi(strArray[1])
	// 返回处理后的结果
	return fmt.Sprintf("%v", a+b)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution0(string(line)))
	}
}
