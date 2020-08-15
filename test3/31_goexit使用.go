package main

import (
	"fmt"
	"runtime"
)

func test(){
	defer fmt.Println("ccccccccccccccc")
	//终止所在的协程
	runtime.Goexit()
	fmt.Println("ddddddddddddddddd")
}
func main(){
	go func(){
		fmt.Println("aaaaaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbbbb")
	}()
	for{
	}
}
