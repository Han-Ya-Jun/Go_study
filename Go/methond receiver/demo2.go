package main

import "fmt"

/*
* @Author:hanyajun
* @Date:2019/8/8 17:57
* @Name:main
* @Function:
 */

type Flags []string

func (f *Flags) Add(flag string) {
	*f = append(*f, flag)
}

func main() {
	var f Flags
	f.Add("a")
	fmt.Println(f)
}
