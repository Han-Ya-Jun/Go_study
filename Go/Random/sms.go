package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/6/10 0:21
* @Name:Random
* @Function:验证码
 */

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
}
