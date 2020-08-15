package main

import "fmt"
/*
	切片和底层
 */
func main(){
	a:=[]int{1,4,6,0,0}
	s:=a[1:4:4]
	fmt.Println("s = ",s)
	fmt.Println("len(s) = ",len(s))
	fmt.Println("cap(s) = ",cap(s))

	fmt.Println("------------------")
	s1:=make([]int,5,10)
	fmt.Println("s = ",s1)
	fmt.Println("len(s) = ",len(s1))
	fmt.Println("cap(s) = ",cap(s1))
	fmt.Println("------------------")

	s2:=[]int{1,3,4,5,6,7,8,9,0,1,42,4,234,2}
	s3:=s2[1:4:10]
	s4:=s3[2:6]
	fmt.Println(s3)
	fmt.Println(s4)

}