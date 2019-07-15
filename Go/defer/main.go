package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/7/1 15:55
* @Name:_defer
* @Function:
 */
func main() {
	fmt.Println(defunc1(1))
	fmt.Println(defunc2(1))
	fmt.Println(defun3(1))
}

func defunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
func defunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func defun3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
