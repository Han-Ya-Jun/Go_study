package main

/*
* @Author:hanyajun
* @Date:18-11-27 下午4:03
* @Name:Go_study/main
* @Function:
 */
func main() {

}

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{-1, -1}
	}
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{j, i}
		} else {
			m[target-v] = i
		}
	}
	return []int{-1, -1}
}
