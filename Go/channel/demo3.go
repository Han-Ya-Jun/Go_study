package main

import (
	"fmt"
	"time"
)

/*
* @Author:hanyajun
* @Date:18-12-10 下午2:01
* @Name:Go_study/main
* @Function:
 */
func main() {
	a := make(chan struct{})
	defer func() {
		if err := recover(); err != nil {
			println("********************")
		}
	}()
	go TestA()
	time.Sleep(8 * time.Second)
	fmt.Println("main over")
	<-a
}

func TestA() {

	panic("")
	chRun := make(chan bool)
	go TestB(chRun)
	select {
	case <-chRun:
		fmt.Println("TestA chRUn")
	case <-time.After(time.Duration(2*1) * time.Second):
		fmt.Println("TestA timeout")
	}
}

func TestB(gLimit chan bool) {
	defer func() {
		fmt.Println("TestB defer")

		fmt.Println("TestB over")

	}()
	//gLimit <- true
	time.Sleep(4 * time.Second)
}
