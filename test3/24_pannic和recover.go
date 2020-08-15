package main

import "fmt"

func testa(){
	fmt.Println("aaaaaaaaaa")
}
func testb(){
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("bbbbbbbbbb")

	a:=[3]int{1,2,3}
	fmt.Println(a)

	panic("this is panic test")//调用pannic主动停止程序
}
func testc(){
	fmt.Println("cccccccccc")
}

func main(){
	testa()
	testb()
	testc()
}
