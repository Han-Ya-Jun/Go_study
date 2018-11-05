package main

/*
* @Author:hanyajun
* @Date:2018/8/8 上午10:43
* @Name:go-study:
* @Function:求最大连续数组的最大和
 */
func main() {
	a := []int{2, -3, 3, 7, 8, 1}
	println(MaxSubSum(a))
}
func MaxSubSum(array []int) int {
	maxSum := 0
	thisSum := 0
	for _, j := range array {
		thisSum += j
		if thisSum > maxSum {
			maxSum = thisSum
			/*如果累加和出现小于0的情况，
			  则和最大的子序列肯定不可能包含前面的元素，
			  这时将累加和置0，从下个元素重新开始累加  */
		} else if thisSum < 0 {
			thisSum = 0
		}

	}
	return maxSum
}
