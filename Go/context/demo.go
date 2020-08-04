package main

import (
	"context"
	"fmt"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/10/22 19:35
* @Name:context
* @Function:
 */
func main() {
	////cancel
	//ctx, cancel := context.WithCancel(context.Background())
	//go work(ctx, "work1")
	//
	//time.Sleep(time.Second * 3)
	//cancel()
	//time.Sleep(time.Second * 1)
	//
	//// with value
	//ctx1, valueCancel := context.WithCancel(context.Background())
	//valueCtx := context.WithValue(ctx1, "key", "test value context")
	//go workWithValue(valueCtx, "value work", "key")
	//time.Sleep(time.Second * 3)
	//valueCancel()

	// timeout
	ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go work(ctx2, "time cancel")

	cancel()

	time.Sleep(time.Second * 10)
	fmt.Println("*************")
	//
	//// deadline
	//ctx3, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	//go work(ctx3, "deadline cancel")
	//time.Sleep(time.Second * 5)
	//deadlineCancel()
	//
	//time.Sleep(time.Second * 3)

}

func workWithValue(ctx context.Context, name string, key string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key))
			println(name, " get message to quit")
			return
		default:
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}
