package main

import "fmt"

/*
* @Author:hanyajun
* @Date:18-11-13 下午4:49
* @Name:Go_study/main
* @Function:
 */
func main() {
	a := []int{5, 10}
	fmt.Println(coinChange2(a, 15))
}
func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 || len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := len(coins) - 1; i >= 0; i-- {
		for j := 0; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = dp[j-coins[i]] + dp[j]
			}
		}

	}
	return dp[amount]
}
