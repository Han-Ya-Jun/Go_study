package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution4(line string) string {
	// 在此处理单行数据
	var maxLen int
	m := make(map[string]string)
	strArray := strings.Split(line, ",")
	for _, v := range strArray {
		m[v] = v
	}

	for _, v := range strArray {
		//寻找上边界
		var len1 int
		a, _ := strconv.Atoi(v)
		for _, ok := m[fmt.Sprintf("%v", a)]; ok; {
			len1++
			delete(m, fmt.Sprintf("%v", a))
			a++
			_, ok = m[fmt.Sprintf("%v", a)]
		}
		//寻找下边界
		b, _ := strconv.Atoi(v)
		b = b - 1
		for _, ok := m[fmt.Sprintf("%v", b)]; ok; {
			len1++
			delete(m, fmt.Sprintf("%v", b))
			b--
			_, ok = m[fmt.Sprintf("%v", b)]
		}
		if len1 > maxLen {
			maxLen = len1
		}
	}
	// 返回处理后的结果
	return fmt.Sprintf("%v", maxLen)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution4(string(line)))
	}
}
