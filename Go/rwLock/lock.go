package main

import (
	"time"
)

/*
* @Author:hanyajun
* @Date:18-11-13 下午9:09
* @Name:Go_study/rwLock
* @Function:
 */
var f bool

func main() {
	c := make(chan struct{})
	f = false

	read(0)
	f = true
	tick := time.NewTicker(time.Second * 1)
	t := 1
	for {
		select {
		case <-tick.C:
			t++
			println(t)
			println("==============================================================================")
			if f {
				f = false
				read(3)
				f = true
			} else {
				println("****************not finished")
			}
		}
	}
	<-c
	//m = new(sync.RWMutex)
	//
	//// 写的时候啥也不能干
	//go read(2)
	//go write(1)
	//
	//go write(3)
	//
	//time.Sleep(2 * time.Second)
}

func read(i int) {
	println("read start")
	time.Sleep(10 * time.Second)
	println("read over")
}
