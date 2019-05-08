package main

import (
	"sync"
	"time"
)

/*
* @Author:15815
* @Date:2019/4/29 10:15
* @Name:main
* @Function:
 */

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Add(1)
		time.Sleep(time.Millisecond)
		wg.Done()
	}()
	wg.Done()
	wg.Wait()
}
