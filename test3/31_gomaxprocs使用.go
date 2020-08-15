package main

import (
	"fmt"
	"runtime"
)

func main(){
	//多核交替频率高，时间片段
	n:=runtime.GOMAXPROCS(6)
	fmt.Println("n =",n)
	for{
		go fmt.Print(1)
		fmt.Print(0)
	}
}
