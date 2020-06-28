package main

import (
	"fmt"
	"reflect"
)

/*
* @Author: HanYaJun
* @Date: 2020/4/16 4:55 下午
* @Name: _interface
* @Function:
 */

type MyError struct {
	str string
}

func (m *MyError) Error() string {
	return m.str
}

func main() {
	err := readDataFromMemory()
	// nil
	if err != nil {
		fmt.Println("err1")
		fmt.Println(err)
		return
	}

	//pt=x12345 vt=0221
	err = daoTest()
	if IsNotNil(err) {
		fmt.Println("err2")
		fmt.Println(err)
		return
	}

	fmt.Println(err)
	fmt.Println("success")
}

func readDataFromMemory() error {
	return nil
}

func daoTest() error {
	return nil
}

func IsNotNil(i interface{}) bool {
	iType := reflect.ValueOf(i)
	if iType.Kind() == reflect.Ptr {
		return iType.IsNil()
	}
	return false
}
