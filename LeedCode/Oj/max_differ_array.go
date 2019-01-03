package main

/*
* @Author:hanyajun
* @Date:2018/8/29 下午8:50
* @Name:go-study:
* @Function:
 */
func main() {
	a := []int{-10, -2, 3, 6, 8, -13}
	println(GetMaxDifferArray(a))
}

func GetMaxDifferArray(a []int) int {
	if len(a) <= 1 && len(a) > 0 {
		return a[0]
	} else if len(a) == 0 {
		return 0
	}
	var minLeft int
	var maxDiffer int
	for i, j := range a {
		if i == 0 {
			maxDiffer = a[1] - a[0]
			if a[1] > a[0] {
				minLeft = a[0]
			} else {
				minLeft = a[1]
			}
		} else if i > 2 {
			if minLeft > j {
				minLeft = j
			}
			if j-minLeft > maxDiffer {
				maxDiffer = j - minLeft
			}
		}
	}
	return maxDiffer
}
