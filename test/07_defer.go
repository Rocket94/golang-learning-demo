package main

import "fmt"

func main(){
	//defer延迟调用，main结束前调用，栈结构
	defer fmt.Println("bbbbbbbbbbbbbb")
	fmt.Println("aaaaaaaaaaaaaa")

	a:=10
	b:=20
	defer func(a,b int){
		fmt.Printf("a=%d,b=%d\n",a,b)
	}(a,b)

	a=111
	b=222
	func(){
		fmt.Printf("a=%d,b=%d\n",a,b)
	}()

}

