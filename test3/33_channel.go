package main

import (
	"fmt"
	"time"
)

//goroutine 通过通信来共享内存，而不是反过来

func Printer(data string){
	for _,d:=range data{
		fmt.Print(string(d))
		time.Sleep(time.Second)
	}
	fmt.Println()
}

var ch=make(chan int)
func Person1(a string)  {
	Printer(a)
	ch<-666//从管道写数据，发送
}
func Person2(a string)  {
	<-ch//如果管道内有数据，接受，如果没有就会阻塞y
	Printer(a)
}
func main(){
	go Person1("hello")
	go Person2("world")
	for{
	}
}
