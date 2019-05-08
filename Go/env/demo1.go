package main

import (
	"fmt"
	"os"
)

/*
* @Author:15815
* @Date:2019/5/8 8:49
* @Name:env
* @Function:获取环境变量的值
 */

func main() {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	fmt.Println("**************************")
	goPath := os.Getenv("GOPATH")
	fmt.Printf("GOPATH is %s\n", goPath)
}
