package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	defer ants.Release()

	runTimes := 1000

	// Use the common pool
	var wg sync.WaitGroup
	//for i := 0; i < runTimes; i++ {
	//	wg.Add(1)
	//	ants.Submit(func() {
	//		demoFunc()
	//		wg.Done()
	//	})
	//}
	//wg.Wait()
	//fmt.Printf("running goroutines: %d\n", ants.Running())
	//fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the size of goroutine pool and 1 second for expired duration
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Serve(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n**************************%d", p.Running(), runtime.NumGoroutine())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
