package main

import (
	"fmt"
	"time"
)

func main()  {
	for {
		go func(){
			fmt.Print(0)
			time.Sleep(100*time.Microsecond)
		}()
		fmt.Print(1)
		time.Sleep(100*time.Microsecond)
	}
}