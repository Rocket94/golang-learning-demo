package main

import "fmt"
//go函数可以返回多个值
func test()(a,b,c int){
	return 7,8,9
}
func main(){
	//变量
	var a,b int
	//初始化
	var c string="abc"
	//自动推导类型并初始化，同一变量名只能用一次
	d:=2.1
	//多重赋值
	a,b=10,20
	fmt.Println(a+b)
	fmt.Println(c)
	fmt.Println(d)
	//匿名变量，丢弃数据不处理，配合返回值使用
	var _,f,_ int=test()
	fmt.Println(f)
	//常量，不能：=
	const a1 int =10
	const a2 =2.1
	const a3,a4=1,4.2
	fmt.Println(a1,a2)
	fmt.Println(a3,a4)
}