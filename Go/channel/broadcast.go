package main

import (
	"fmt"
	"sync"
)

/*
* @Author:hanyajun
* @Date:2019/5/31 10:03
* @Name:main
* @Function: go广播
 */

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range ch {
			fmt.Println("A", v)
		}
		wg.Done()
	}()
	go func() {
		for v := range ch {
			fmt.Println("B", v)
		}
		wg.Done()
	}()

	ch <- 1
	ch <- 2
	close(ch)
	wg.Wait()
}
