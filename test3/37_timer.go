package main

import (
	"fmt"
	"time"
)

func main(){
	timer:=time.NewTimer(2*time.Second)
	fmt.Println("当前时间：",time.Now())

	t:=<-timer.C
	fmt.Println("当前时间：",t)

	t2:=<-time.After(2*time.Second)
	fmt.Println("当前时间：",t2)

	time.Sleep(2*time.Second)
	fmt.Println("时间到",time.Now())
}
