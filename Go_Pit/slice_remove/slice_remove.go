package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2018/3/29 下午1:25
* @Name:Go_study:
* @Function:slice 去除某个元素（元素不能有重复）
 */
func main() {
	var s = []string{"ca", "ab", "ec"}
	for k, v := range s {
		if v == "ab" {
			kk := k + 1
			s = append(s[:k], s[kk:]...)
		}
	}
	fmt.Println(s)
}
