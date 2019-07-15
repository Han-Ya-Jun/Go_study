package main

import (
	"encoding/base64"
	"fmt"
)

/*
* @Author:hanyajun
* @Date:2019/6/24 16:30
* @Name:base64
* @Function:
 */

func main() {

	var str = "ZHONGGUOnihao123"
	strbytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	decodestr := string(decoded)
	fmt.Println(decodestr, err)

}
