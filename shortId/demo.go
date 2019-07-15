package main

import (
	"fmt"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/teris-io/shortid"
)

/*
* @Author:hanyajun
* @Date:2019/6/24 17:15
* @Name:shortId
* @Function:
 */
func main() {
	fmt.Println(shortid.Generate())
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newId, _ := currWoker.NextId()
	println(newId)
}
