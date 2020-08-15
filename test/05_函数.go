package main

import (
	"fmt"
)

//不定参数类型
func MyFunc1(args...int){
	MyFunc2(args...)
}
func MyFunc2(tmp...int){
	for _,data:=range tmp{
		fmt.Println("tmp=",data)
	}
}
//切片
func MyFunc3(args...int)  {
	MyFunc2(args[2:]...)
}

func main(){
	MyFunc3(1,3,4,5,5,5)
	fmt.Println(MyFunc4())
	fmt.Println(MyFunc5(5))
	//起别名
	type FuncType func(int)(int)
	var fTest FuncType
	fTest=MyFunc5
	fmt.Println(fTest(4))

	//回调
	a:=Calc(1,1,Sub)
	fmt.Println(a)
}

func MyFunc4()(a ,b int){
	a,b=1,2
	return
}
//递归
func MyFunc5(n int)(result int){
	if n==0{
		return 1
	}
	if n==1{
		return 1
	}
	return MyFunc5(n-1)+MyFunc5(n-2)

}
//回调函数,函数有一个参数是类型函数,实现多态
type FuncType func(a,b int)int

func Calc(a,b int,fTest FuncType)(result int){
	fmt.Println("Calc")
	result=fTest(a,b)
	return
}

func Add(a,b int)int{
	return a+b
}
func Sub(a,b int)int{
	return a-b
}
