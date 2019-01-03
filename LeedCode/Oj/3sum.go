package main

import (
	"fmt"
	"sort"
)

/*
* @Author:hanyajun
* @Date:18-11-27 下午3:48
* @Name:Go_study/main
* @Function:
 */
func main() {

	var nums = []int{-1, 0, 1, 2, -1, -4}
	three := threeSum(nums)
	fmt.Println(three)
}

func threeSum(nums []int) [][]int {
	if len(nums) <= 0 {
		return [][]int{}
	}

	if !sort.IntsAreSorted(nums) {
		sort.Ints(nums)
	}

	var sum int
	var res [][]int
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		front := i + 1
		back := len(nums) - 1

		for front < back {
			sum = nums[front] + nums[back]

			if sum < target {
				front++
			} else if sum > target {
				back--
			} else {
				var tripliet = []int{nums[i], nums[front], nums[back]}
				res = append(res, tripliet)

				for front < back && nums[front] == tripliet[1] {
					front++
				}
				for front < back && nums[back] == tripliet[2] {
					back--
				}
			}
		}

		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}

	return res
}
