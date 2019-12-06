package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
* @Author:hanyajun
* @Date:2019/7/29 14:20
* @Name:http
* @Function:
 */

type DockFindAddressRequestParam struct {
	SrcId  int64 `json:"src_id,string"`
	DestId int64 `json:"dest_id,string"`
}
type ResponseData struct {
	Code int `json:"code"`
	Msg  int `json:"msg"`
	Data []*DockFindResultResponse
}
type DockFindResultResponse struct {
	SrcId  int64 `json:"src_id,string"`
	DestId int64 `json:"dest_id,string"`
	DockId int64 `json:"dock_id,string"`
}

func main() {
	param := []*DockFindAddressRequestParam{&DockFindAddressRequestParam{
		SrcId:  72059260488253440,
		DestId: 72059260489203712,
	}}
	b, _ := json.Marshal(param)
	fmt.Println(string(b))
	res, _ := http.Post("http://localhost:8000/resource-manage-svc/api/v1/address/get_dock/72059260489203712", "application/json", bytes.NewBuffer(b))
	result, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", result)
	var resultList ResponseData

	err := json.Unmarshal(result, &resultList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", resultList.Data)
}
