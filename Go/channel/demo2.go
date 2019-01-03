package main

import (
	"fmt"
	"time"
)

/*
* @Author:hanyajun
* @Date:18-12-4 下午3:25
* @Name:Go_study/main
* @Function:
 */
func main() {
	glimit := make(chan bool, 10)
	g := make(chan struct{})
	for i := 0; i < 100; i++ {
		glimit <- true
		go RunSkuTask(i, 2, glimit)
	}
	<-g
}

func RunSkuTask(productIds int, days int, gLimit chan bool) {
	chRun := make(chan bool)
	go CacheGoodsSaleInfo(productIds, days, chRun)
	select {
	case <-chRun:
		<-gLimit
	case <-time.After(time.Duration(60*1) * time.Second):
		re := fmt.Sprintf("sku task id %v , timeout", productIds)
		println(re)
		<-gLimit
	}
	println("123123123123123123")
}

func CacheGoodsSaleInfo(productIds int, days int, gLimit chan bool) {
	println("*****************************")
	gLimit <- true

}
