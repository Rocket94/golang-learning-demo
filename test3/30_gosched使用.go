package main

import (
	"fmt"
	"runtime"
)

func main(){
	go func(){
		for i:=0;i<5;i++{
			fmt.Println("go")
		}
	}()
	//让出时间片，先让别的协程执行，之后再回来执行此协程
	for i:=0;i<2;i++{
		runtime.Gosched()
		fmt.Println("hello")
	}

}
