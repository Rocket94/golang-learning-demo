package main

import (
	"fmt"
)

func main() {
	ch:=make(chan int)

	go func() {
		for i:=0;i<5;i++{
			ch<-i
		}
		close(ch)
	}()

	for{
		//如果OK是true，说明管道没有关，管道关了，无法发数据，可以接数据
		if num,ok:=<-ch;ok {
			fmt.Println("num =",num)
		}else {
			break
		}
	}
}
