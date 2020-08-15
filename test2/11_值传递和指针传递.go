package main

import "fmt"

func main(){
	a,b:=1,2
	swap1(a,b)
	fmt.Printf("a = %d,b = %d\n",a,b)
	swap2(&a,&b)
	fmt.Printf("a = %d,b = %d\n",a,b)
	c:=[]int{1,3,5,7,9}
	chSlice(c)
	fmt.Println(c)
}
//变量，结构体，数组是值传递
func swap1(a,b int){
	a,b=b,a
}
func swap2(p,q *int){
	*p,*q=*q,*p
}
//切片和map是引用传递
func chSlice(a[]int){
	a[0],a[1]=a[1],a[0]
}

