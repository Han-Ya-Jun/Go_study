package main

/*
* @Author:hanyajun
* @Date:18-11-8 下午8:19
* @Name:Go_study/main
* @Function:合并两个有序数组
 */
func main() {

}
func merge(nums1 []int, m int, nums2 []int, n int) {
	length := m + n
	for m > 0 && n > 0 {
		length--
		if nums1[m-1] > nums2[n-1] {
			m--
			nums1[length] = nums1[m]

		} else {
			n--
			nums1[length] = nums2[n]
		}
	}

	for m > 0 {
		length--
		m--
		nums1[length] = nums1[m]
	}
	for n > 0 {
		length--
		n--
		nums1[length] = nums2[n]
	}

}
