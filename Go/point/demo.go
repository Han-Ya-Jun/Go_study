package main

import (
	"fmt"
	"time"
)

/*
* @Author: HanYaJun
* @Date: 2020/5/12 11:58 上午
* @Name: point
* @Function:
 */

type A struct {
	b *B
}

type B struct {
	c *string
}

func main() {
	str := "s"
	b :=
		&B{
			c: &str,
		}
	a := A{
		b: b,
	}
	var f = func() {}
	f = func() {
		c := "ddd"
		a.b = &B{
			c: &c,
		}
		d := a.b
		a.b = b
		if !(*d.c == "s" || *d.c == "ddd") {
			fmt.Println(*d.c)
			panic(*d.c)
		}
		fmt.Print("%v", a)
		for i := 0; i < 1000; i++ {
			go f()
		}
	}
	for i := 0; i < 10000*10000; i++ {
		go f()
	}
	time.Sleep(time.Second * 10000)
}
