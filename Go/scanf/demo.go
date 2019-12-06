package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/11/20 22:34
* @Name:scanf
* @Function:
 */

func culculate(num []float64, m string) float64 {

	switch m {
	case "+":
		return num[0] + num[1]
	case "-":
		return num[0] - num[1]

	case "*":
		return num[0] * num[1]

	case "/":
		return num[0] / num[1]
		//不能一个case一个return
	}
	return num[0]
}

//%%
func main() {

	var x float64
	var sign string
	var lastSign string
	var num []float64
	var str string
	var result float64
	for {
		fmt.Println("input e your numbers and sign")
		fmt.Scanln(&x, &sign)
		num = append(num, x)
		if sign == "=" {
			result = culculate(num, lastSign)
			fmt.Println(result)
			break
		}
		str += fmt.Sprintf("%v %v ", x, sign)
		fmt.Println(str)
		if len(num) == 2 {
			result = culculate(num, sign)
			num = []float64{result}
			lastSign = sign
		}

	}

}
