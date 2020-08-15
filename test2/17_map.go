package main

import "fmt"

func main(){
	//创建
	m1:=map[int]string{1:"彭荣坤",2:"王慧咏"}
	fmt.Println("m1 =",m1)
	//修改
	m2:=m1
	m2[1]="nihao"
	m2[2]="wohao"
	fmt.Println("m2 =",m2)
	//遍历
	m3:=map[int]string{1:"彭荣坤",2:"王慧咏",3:"王大爷",5:"小彭"}
	for k,v:=range m3{
		fmt.Println("key and value is",k,v)
	}
	//判断
	value,ok:=m3[55]
	if ok{
		fmt.Println("key exists,value is",value)
	}else {
		fmt.Println("dont have the key")
	}
	//删除
	delete(m3,3)
	for k,v:=range m3{
		fmt.Println("key and value is",k,v)
	}


}
