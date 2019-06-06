package main

import (
	"fmt"
	"sync"
)

/*
* @Author:hanyajun
* @Date:2019/5/31 14:11
* @Name:main
* @Function:
 */

func main() {
	chA := make(chan int)
	chB := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range chA {
			fmt.Println("A", v)
		}
		wg.Done()
	}()

	go func() {
		for v := range chB {
			fmt.Println("B", v)
		}
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		chA <- i
		chB <- i
	}
	close(chA)
	close(chB)
	wg.Wait()
}
