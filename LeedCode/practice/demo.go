package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/*
* @Author: HanYaJun
* @Date: 2020/4/23 5:46 下午
* @Name: practice
* @Function:
 */

func main() {
	m := map[string]string{"1": "bar", "2": "foo.bar", "3": "foo.foo", "4": "baz.cloudmall.com", "5": "baz.cloudmall.ai"}
	//{
	//  "bar": "1",
	//  "foo": {
	//    "bar": "2",
	//    "foo": "3"
	//  },
	//  "baz": {
	//    "cloudmall": {
	//      "com": "4",
	//      "ai": "5"
	//    }
	//  }
	//}
	aa := 1231231231231
	fmt.Println(strconv.FormatUint(uint64(aa), 2))
	n := make(map[string]interface{})
	for k, v := range m {
		kList := strings.Split(v, ".")

		for i, vv := range kList {
			if i == len(kList)-1 {
				n[vv] = k
				break
			}
			if value, ok := n[vv]; !ok {
				vm := map[string]interface{}{kList[i+1]: k}
				n[vv] = vm
			} else {
				if vvvv, ok := value.(map[string]interface{}); ok {
					vvvv[vv] = map[string]interface{}{kList[i+1]: k}
					n[vv] = vvvv
				}
			}
		}
	}
	b, _ := json.Marshal(n)
	fmt.Printf("%s", b)
}
