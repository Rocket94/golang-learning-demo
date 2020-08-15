package main

import (
	"fmt"
	"time"
)
func newTask(){
	for{
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}
//并发：单个CPU运行，时间片轮转；并行，多个CPU运行
//主协程退出，子携程也退出
func main(){
	go newTask()//新建一个协程
	for{
		fmt.Println("this is a goroutine")
		time.Sleep(time.Second)
	}

}
