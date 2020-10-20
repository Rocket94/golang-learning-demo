package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main(){
	//注册回调函数
	http.HandleFunc("/hello", handler)

	//绑定tcp监听地址，并开始接受请求，然后调用服务端处理程序来处理传入的连接请求。
	//params：第一个参数 addr 即监听地址；第二个参数表示服务端处理程序，通常为nil
	//当参2为nil时，服务端调用 http.DefaultServeMux 进行处理
	http.ListenAndServe("127.0.0.1:2222", nil)
}
type Report struct {
	RequestTime uint64         `json:"requestTime"` // 请求总时间
	MaxTime     uint64         `json:"maxTime"`     // 最大时长
	MinTime     uint64         `json:"minTime"`     // 最小时长
	SuccessNum  uint64         `json:"successNum"`  // 成功请求数
	FailureNum  uint64         `json:"failureNum"`  // 失败请求数
	NewOrderNum uint64         `json:"newOrderNum"` // tpmc
	ConCurrency uint64         `json:"conCurrency"` // 并发数
	ErrCode     map[int]int    `json:"errCode"`     // 错误码/错误个数
	ErrCodeMsg  map[int]string `json:"errCodeMsg"`  // 错误码描述
	Status      bool           `json:"status"`
}
func handler (w http.ResponseWriter, r *http.Request){
	fmt.Println(r.RemoteAddr, "连接成功")  	//客户端网络地址
	w.Header().Set("Content-Type", "application/json")
	post := &Report{
		RequestTime: 0,
		MaxTime:     0,
		MinTime:     0,
		SuccessNum:  0,
		FailureNum:  0,
		NewOrderNum: 0,
		ConCurrency: 0,
		ErrCode:     nil,
		ErrCodeMsg:  nil,
		Status:      false,
	}
	json, err := json.Marshal(post)
	if err != nil {
		return
	}
	w.Write(json)
}
