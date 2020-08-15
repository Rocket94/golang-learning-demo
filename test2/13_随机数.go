package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//设置种子，只需一次
	rand.Seed(time.Now().UnixNano())

	for i:=0;i<5;i++{
		fmt.Printf("%d\n",rand.Intn(100))
	}
}
