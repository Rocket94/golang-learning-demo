package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"daming"`
	age int
	IsMarried bool
	sex byte
}

func main(){
	s:=Person{"彭荣坤",26,false,'m'}
	//编码，根据内容生成json
	buf,err:=json.Marshal(s)
	if err!=nil{
		fmt.Println("err is",err)
		return
	}
	fmt.Println("buf =",string(buf))
	fmt.Println("--------------------")
	//map存储json
	m:=make(map[string]interface{},4)
	m["name"]="王慧咏"
	m["hobby"]=[]string{"唱","跳","RAP","篮球"}
	m["isMarried"]=false
	m["age"]=18.888
	result,err2:=json.Marshal(m)
	if err2!=nil{
		fmt.Println("err =",err)
		return
	}
	fmt.Println("result =",string(result))
	fmt.Println("--------------------")
	//解析json
	var tmp Person
	err3:=json.Unmarshal(buf,&tmp)//地址传递
	if err3!=nil{
		fmt.Println("err =",err3)
		return
	}
	fmt.Printf("tmp = %v\n",tmp)
}