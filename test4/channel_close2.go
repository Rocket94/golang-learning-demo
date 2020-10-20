package main

import (
	"fmt"
	"sync"
	"time"
)

type abc struct {
	num int
	name string
}

func main(){
	ch:=make(chan abc,20000)
	close(ch)
	var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			for j:=1;j<3000;j++{
				a:=abc{
					num:  j,
					name: "a",
				}
				ch<-a
			}
			wg.Done()
		}()

	go func() {
		for data:=range ch{
			time.Sleep(100*time.Microsecond)
			fmt.Println(data)
			//if j==nil{
			//	return
			//}

		}
		fmt.Println("kid over")
	}()
	wg.Wait()
	time.Sleep(1*time.Second)
	fmt.Println("over")
	for{

	}

}
