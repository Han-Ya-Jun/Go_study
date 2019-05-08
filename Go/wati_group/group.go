package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* @Author:hanyajun
* @Date:2018/8/11 下午2:39
* @Name:go-study:
* @Function:
 */

//func test(shownum int, waitgroup *sync.WaitGroup) {
//	fmt.Println(shownum)
//	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
//}

func main() {
	//var waitgroup sync.WaitGroup
	b := make(chan string, 100)
	go CreateTask(b)
	go Run(b)
	//waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	fmt.Println("done!")
	time.Sleep(time.Hour)
}

func CreateTask(c chan string) {
	var waitgroup sync.WaitGroup
	println("*************")
	a := make(chan bool, 10000)
	for i := 0; i < 10; i++ {
		println(i)
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go CreateSubTask(i, a, &waitgroup)
	}
	waitgroup.Wait()
	println("****task is done")
}
func CreateSubTask(n int, a chan bool, w *sync.WaitGroup) {
	s := fmt.Sprintf("************%v", n)
	println(s)
	w.Done()
	a <- true
}
func Run(a chan string) {
	for {
		select {
		case plan := <-a:
			println(plan)

		}
	}
}
