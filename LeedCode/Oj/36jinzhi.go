package main

import "fmt"

/*
* @Author:hanyajun
* @Date:18-11-22 下午3:51
* @Name:Go_study/main
* @Function:36进制加法
 list_36 = ["0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]

*/
func main() {
	fmt.Println(Add("1B", "2X"))
}
func Add(str1 string, str2 string) string {
	List36 := []string{
		"0", "1", "2", "3", "4", "5", "6",
		"7", "8", "9", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	i := len(str1) - 1
	j := len(str2) - 1
	var sum string
	var tem int //进位
	for i >= 0 && j >= 0 {
		s := GetInt(str1[i]) + GetInt(str2[j]) + tem
		if s >= 36 {
			tem = 1
			sum = List36[s%36] + sum
		} else {
			tem = 0
			sum = List36[s] + sum
		}
		i--
		j--
	}
	for i >= 0 {
		s := GetInt(str1[i]) + tem
		if s >= 36 {
			tem = 1
			sum = List36[s%36] + sum
		} else {
			tem = 0
			sum = List36[s] + sum
		}
		i--
	}
	for j >= 0 {
		s := GetInt(str2[i]) + tem
		if s >= 36 {
			tem = 1
			sum = List36[s%36] + sum
		} else {
			tem = 0
			sum = List36[s] + sum
		}
		j--
	}
	if tem != 0 {
		sum = "1" + sum
	}
	return sum
}

//将'0'-'9'映射到数字0-9，将'a'-'z'映射到数字10-35
func GetInt(a uint8) int {
	if a-'0' > 0 && a <= '9' {
		return int(a - '0')
	} else {
		return int(a-'A') + 10
	}

}
