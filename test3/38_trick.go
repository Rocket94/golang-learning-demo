package main

import (
	"fmt"
	"time"
)

func main(){
	t:=time.NewTicker(time.Second)
	for i:=0;i<100;i++{
		<-t.C
		fmt.Println("i =",i)
	}
}
