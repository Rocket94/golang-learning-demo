package main

import "fmt"

func main(){
	var id [50]int
	for i,_:=range id{
		id[i]=i
		fmt.Println(id[i])
	}
	fmt.Printf("长度为%d\n", len(id))

	b:=[5]int{2:3,4:5}
	for _,data:=range b{
		fmt.Println(data)
	}
	fmt.Println(b)
}
