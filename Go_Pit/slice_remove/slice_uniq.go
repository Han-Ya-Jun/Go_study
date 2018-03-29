package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2018/3/29 下午1:29
* @Name:Go_study: 
* @Function:
 */
func main() {
	var s = []string{"ca", "ab", "ec", "ca", "ab", "ec", "ca", "ab", "ec"}
	index := 0
	endIndex := len(s) - 1
	var result = make([]string, 0)
	for k, v := range s {
		if v == "ab" {
			result = append(result, s[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, s[index:endIndex+1]...)
		}
	}
	fmt.Println(result)
}
