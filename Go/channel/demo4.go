package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

var sum2 int32

func myFunc2(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum2, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc2() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

type Param struct {
}

func main() {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool
	var wg sync.WaitGroup
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants.Submit(func() error {
			demoFunc2()
			wg.Done()
			return
		})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the size of goroutine pool and 1 second for expired duration
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc2(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks
	ps := &Param{}
	for i := 0; i < runTimes; i++ {
		wg.Add(i)
		p.Serve(ps)
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum2)
}
