package main

import "fmt"

/*
* @Author: HanYaJun
* @Date: 2019/12/6 3:54 下午
* @Name: bit
* @Function:
 */

func main() {
	fmt.Println(GetTransportationSpecialityMask(1, 8))
}

func GetTransportationSpecialityMask(start, end uint) uint64 {
	var mask uint64
	for i := start; i <= end; i++ {
		mask += 1 << (i - 1)
		fmt.Println(mask)
	}
	return mask
}
