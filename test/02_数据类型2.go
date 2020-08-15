package main

import (
	"fmt"
)

func main(){
	//iota常量自动生成器，每隔一行自动累加1
	//iota给常量复制用
	const (
		a=iota
		b=iota
		c=iota
	)
	fmt.Printf("a=%d,b=%d,c=%d\n", a,b,c)
	//遇见const重置为零
	const d=iota
	fmt.Printf("d=%d\n",d)
	//同一行都是一样的值

	var e bool=true
	if(e){
		print("对")
	}
	//%T输出数据类型
	var f1 float32=3.14
	f2:=3.14
	fmt.Println("f1=",f1)
	fmt.Printf("f2=%T\n",f2)
	//%c字符方式打印,%d整型打印
	var ch byte
	ch=97
	fmt.Printf("%c,%d\n",ch,ch)

	ch='a'
}
