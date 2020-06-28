package main

import (
	"fmt"
)

/*
* @Author:hanyajun
* @Date:2019/7/23 20:32
* @Name:slince
* @Function:
 */

func main() {
	s1 := make([]int, 0, 100)
	s2 := s1
	s1 = append(s1, 1)
	fmt.Println(s2)
	s2 = append(s2, 2)
	fmt.Println(s1)
	array := []int{2, 4}
	ChangeSlince(array)
	fmt.Println(array)

}
func ChangeSlince(s []int) {
	s[0] = 24
}
