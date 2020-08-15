package main

import "fmt"

func main(){
	a:=10
	str:="mario"
	//匿名函数，没有名字定义，没有被调用
	f1:=func(){
		fmt.Println(a,str)
	}
	f1()
	type FuncType func()
	var f2 FuncType
	f2=f1
	f2()
	//闭包改变外面变量
	func(){
		a=666
		str="jack"
		fmt.Println(a,str)
	}()
	fmt.Println(a,str)

	fmt.Println("----------------------")
	fmt.Println(MyFun1())
	fmt.Println(MyFun1())
	fmt.Println(MyFun1())
	fmt.Println(MyFun1())
	fmt.Println("----------------------")
	f:=MyFun2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
func MyFun1()int{
	x:=0
	x++
	return x*x
}
func MyFun2()func()int{
	x:=0
	return func()int{
		x++
		return x*x
	}
}
