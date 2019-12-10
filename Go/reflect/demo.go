package main

import (
	"encoding/json"
	"fmt"
)

/*
* @Author:hanyajun
* @Date:2019/8/16 15:30
* @Name:reflect
* @Function:
 */
type Demo struct {
	Age int
}

type Demo2 struct {
	Name string
}

type Consumers struct {
	Consumer []chan interface{}
}

func main() {
	a := make(chan Demo, 1000)
	b := make(chan interface{}, 1000)
	var chans []chan interface{}
	chans = append(chans, a, b)
	c := Consumers{
		Consumer: chans,
	}

	demo := Demo{Age: 124}
	demo2 := Demo2{
		Name: "sdf",
	}
	demostr, _ := json.Marshal(demo)
	demo2str, _ := json.Marshal(demo2)
	chans[0] <- demostr
	chans[1] <- demo2str
	for _, cc := range c.Consumer {
		fmt.Printf("%s", <-cc)
	}

}
