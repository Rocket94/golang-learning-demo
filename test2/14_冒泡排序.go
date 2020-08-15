package main

import (
	"fmt"
)

func main(){
	a:=[]int{55,88,13,74,99,11,4,5,32,88}
	bubbling(a)
	fmt.Println(a)
}

func bubbling(a []int){
	for i:=0;i<len(a)-1;i++{
		for j:=0;j<len(a)-i-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}

	}

}
