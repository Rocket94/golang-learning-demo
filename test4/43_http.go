package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		ok:=rand.Intn(10)
		if ok==0{
			fmt.Println("yes!",ok)

		}else {
			fmt.Println("oh! no!",ok)
		}

	}
}
