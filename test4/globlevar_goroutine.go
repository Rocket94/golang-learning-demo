package main

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {
	isStop:=false
	for i:=0;i<10;i++{
		go func(num int) {
			for{
				if !isStop{
					fmt.Println(num,":child goroutine is running")
					time.Sleep(1*time.Second)
				}else{
					break
				}
			}
		}(i)
	}

	go func() {
		t := time.After(5 * time.Second)
		<-t
		isStop=true
		fmt.Println("stop turns true")
	}()
		fmt.Println(runtime.NumGoroutine())

	for {
		if isStop{
			fmt.Println(runtime.NumGoroutine())
		}
		time.Sleep(10*time.Second)

	}
}