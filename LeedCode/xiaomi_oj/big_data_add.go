package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution3(line string) string {
	//// 在此处理单行数据
	if strings.Contains(line, " ") {
		line = strings.Replace(line, " ", "", -1)
	}
	strArray := strings.Split(line, "-")
	println(len(strArray[0]) - len(strArray[1]))
	str1, str2 := HandleSame(strArray[0], strArray[1])
	var r string
	var y int
	var result string
	for i := len(str1); i > 0; i-- {
		strA := string([]byte(str1)[i-1 : i])
		strB := string([]byte(str2)[i-1 : i])
		y, r = Reduce(strA, strB, y)
		result = r + result
	}
	return result
}

//考虑两个大数的长度不一样
func HandleSame(str1 string, str2 string) (string, string) {
	len := len(str1) - len(str2)
	if len == 0 {
		return str1, str2
	} else {
		for i := 0; i < len; i++ {
			str2 = "0" + str2
		}
	}
	return str1, str2
}

//每位相加减，考虑借位
func Reduce(a string, b string, c int) (d int, e string) {
	s1, _ := strconv.Atoi(a)
	s2, _ := strconv.Atoi(b)
	r := s1 - s2 + c
	if r < 0 {
		r = r + 10
		d = -1
	}
	return d, fmt.Sprintf("%v", r)
}
func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution3(string(line)))
	}
}
