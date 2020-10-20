package main

import (
	"fmt"
	"runtime"
)

func main()  {

	for i:=0;i<10;i++{
		go func() {
			go func() {
				for{
					fmt.Println(i,"kid : ",runtime.NumGoroutine())
				}
			}()

			for{
				fmt.Println(i," : ",runtime.NumGoroutine())
			}

		}()
	}

	for{
		fmt.Println("main : ",runtime.NumGoroutine())
	}
}