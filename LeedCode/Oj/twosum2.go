package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
* @Author:hanyajun
* @Date:18-11-27 下午4:04
* @Name:Go_study/main
* @Function:
 */
func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(600))
	num := []int{2, 4, 1, 5, 3, 3, 4, 2}
	fmt.Printf("%v", twoSum2(num, 6))
}
func twoSum2(nums []int, target int) [][]int {
	if len(nums) <= 1 {
		return [][]int{}
	}
	var result [][]int
	m := make(map[int]int, len(nums))
	re := make(map[string]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			key := fmt.Sprintf("1:%v2:%v", v, target-v)
			if _, ok := re[key]; !ok {

			}
			r := []int{v, target - v}
			result = append(result, r)
		} else {
			m[target-v] = v
		}
	}
	return result
}
