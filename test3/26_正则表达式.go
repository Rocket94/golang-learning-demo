package main

import (
	"fmt"
	"regexp"
)

func main(){
	buf := "abc azc aac 888 a9c tac"
	//解释规则，解析正则表达式，返回解释器
	reg1 :=regexp.MustCompile("a.c")
	if reg1==nil{
		fmt.Println("regex error")
		return
	}
	//根据规则提取信息
	result1:=reg1.FindAllStringSubmatch(buf,-1)
	fmt.Println("result1 =",result1)

}