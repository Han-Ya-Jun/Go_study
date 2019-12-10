package main

import (
	"fmt"
	"regexp"
)

/*
* @Author:hanyajun
* @Date:2019/7/15 18:23
* @Name:hex
* @Function:
 */

func main() {
	//var a int64
	s := fmt.Sprintf("%.4x", 255)
	fmt.Println(s)
	matched, _ := regexp.MatchString(`[\w-]+@[\w]+(?:\.[\w]+)+`, sourceStr)
	a := "123456"
	fmt.Println(Replace(a, 2, 4, "65"))

}

func Replace(old string, start, end int, new string) string {
	return substr(old, 0, start) + new + substr(old, end, len([]rune(old)))
}

func substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
