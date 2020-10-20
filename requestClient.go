package main

import (
"bytes"
"encoding/json"
"io"
"io/ioutil"
"net/http"
"time"
	"fmt"
)
type Request struct {
	// 请求赋值
	Url         string `json:"url"`        // 请求域名
	ConCurrency uint64 `json:"conCurrent"` // 并发数
	Duration    uint64 `json:"duration"`   // 持续时间（秒）
	Preheating  uint64 `json:"preheating"` // 预热时间（秒）
	Warehouse   uint64 `json:"warehouse"`
}

func main(){
	r:=Request{
		//Url:         "http://10.110.68.25:1501/sut",
		Url: "http://127.0.0.1:2222/hello",
		ConCurrency: 100,
		Duration:    600,
		Preheating:  0,
		Warehouse:   1000,
	}
	s1:=Post("http://127.0.0.1:1500/start",r,"application/json")
	//s1:=Post("http://10.110.68.25:1502/start",r,"application/json")
	//s2:=Post("http://10.110.68.25:1503/start",r,"application/json")
	//s3:=Post("http://10.110.68.25:1504/start",r,"application/json")
	//s4:=Post("http://10.110.68.25:1505/start",r,"application/json")
	//s5:=Post("http://10.110.68.25:1506/start",r,"application/json")
	fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(s3)
	//fmt.Println(s4)
	//fmt.Println(s5)
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
