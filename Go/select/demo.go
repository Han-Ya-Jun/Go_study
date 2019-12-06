package main

import "time"

/*
* @Author:hanyajun
* @Date:2019/8/9 17:08
* @Name:_select
* @Function:
 */

func main() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")

		case <-ch:
			println("case2")
		default:
			println("****************")
		}
	}
}
