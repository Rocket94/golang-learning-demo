package main

import "fmt"

func main(){
	var ch=make(chan string)
	defer fmt.Println("主携程也完事")
	go func() {for i:=0;i<2;i++{
		fmt.Println(i)}
		ch<-"子协程完事"
	}()
	b:=<-ch
	fmt.Println(b)
}
