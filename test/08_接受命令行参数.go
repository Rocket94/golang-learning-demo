package main

import (
	"fmt"
	"os"
)

func main(){
	//接受用户传递的参数,字符串方式
	list:=os.Args

	n:=len(list)
	fmt.Println(n)
	for i:=0;i<len(list);i++{
		fmt.Printf("list[%d]=%s\n",i,list[i])
	}
}
