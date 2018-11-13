package main

import (
	"fmt"
	"math"
)

/*
* @Author:hanyajun
* @Date:18-11-13 下午1:35
* @Name:Go_study/main
* @Function:硬币交换
 */
func main() {
	a := []int{5}
	fmt.Println(coinChange(a, 15))
}

//通过动态规划来解决
func coinChange(coins []int, amount int) int {
	if amount < 0 || len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		min := math.MaxInt64
		for _, c := range coins {
			j := i - c
			if j >= 0 && dp[j] != -1 {
				min = int(math.Min(float64(min), float64(dp[j])))
			}
			if min == math.MaxInt64 {
				dp[i] = -1
			} else {
				dp[i] = min + 1
			}

		}

	}
	return dp[amount]
}
