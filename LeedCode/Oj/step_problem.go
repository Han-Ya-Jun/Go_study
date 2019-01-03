package main

/*
* @Author:hanyajun
* @Date:18-11-21 下午1:17
* @Name:Go_study/main
* @Function:一个环，有n个点, 问从0点出发，经过k步回到原点有多少种方法: d(k, j) = d(k-1, (j-1+n)%n) + d(k-1, (j+1)%n);
 */
func main() {
	println(GetSteps(10, 1))
}

func GetSteps(n int, k int) int {
	if n == 0 {
		return 1
	}
	if n == 2 {
		if k%2 == 0 {
			return 1
		} else {
			return 0
		}
	}
	var dp [100][100]int

	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[0][i] = 0
	}
	for j := 1; j <= k; j++ {
		for i := 0; i < n; i++ {
			dp[j][i] = dp[j-1][(i-1+n)%n] + dp[j-1][(i+1)%n]
		}
	}
	return dp[k][0]
}
